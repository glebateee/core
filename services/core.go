package services

import (
	"context"
	"fmt"
	"reflect"
)

type BindingMap struct {
	factoryFunc reflect.Value
	lifecycle
}

var services = make(map[reflect.Type]BindingMap)

func addService(life lifecycle, factoryFunc any) error {
	factoryFuncType := reflect.TypeOf(factoryFunc)
	if factoryFuncType.Kind() == reflect.Func && factoryFuncType.NumOut() == 1 {
		services[factoryFuncType.Out(0)] = BindingMap{
			factoryFunc: reflect.ValueOf(factoryFunc),
			lifecycle:   life,
		}
		return nil
	}
	return fmt.Errorf("type cannot be used as service: %v", factoryFuncType.Name())
}

var contextReference = (*context.Context)(nil)
var contextReferenceType = reflect.TypeOf(contextReference).Elem()

func resolveScopedService(ctx context.Context, ptr reflect.Value, binding BindingMap) {
	sMap, ok := ctx.Value(serviceKey).(ServiceMap)
	if !ok {
		ptr.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
	} else {
		serviceVal, ok := sMap[ptr.Type()]
		if !ok {
			serviceVal = invokeFunction(ctx, binding.factoryFunc)[0]
			sMap[ptr.Type()] = serviceVal
		}
		ptr.Elem().Set(serviceVal)
	}
}

func resolveServiceFromValue(ctx context.Context, ptr reflect.Value) error {
	serviceType := ptr.Elem().Type()
	if serviceType == contextReferenceType {
		ptr.Elem().Set(reflect.ValueOf(ctx))
		return nil
	} else {
		binding, ok := services[serviceType]
		if ok {
			if binding.lifecycle == Scoped {
				resolveScopedService(ctx, ptr, binding)
			} else {
				ptr.Elem().Set(invokeFunction(ctx, binding.factoryFunc)[0])
			}
			return nil
		}
	}
	fmt.Printf("----------------------%s\n", serviceType.Name())
	return fmt.Errorf("cannot find service: %v", serviceType.Name())

}

func resolveFunctionArguments(ctx context.Context, f reflect.Value, notSevices ...any) []reflect.Value {
	params := make([]reflect.Value, f.Type().NumIn())
	i := 0
	for ; i < len(notSevices); i++ {
		params[i] = reflect.ValueOf(notSevices[i])
	}
	for ; i < len(params); i++ {
		pType := f.Type().In(i)
		pVal := reflect.New(pType)
		if err := resolveServiceFromValue(ctx, pVal); err != nil {
			panic(err)
		}
		params[i] = pVal.Elem()
	}
	return params
}

func invokeFunction(ctx context.Context, f reflect.Value, notSevices ...any) []reflect.Value {
	return f.Call(resolveFunctionArguments(ctx, f, notSevices...))
}
