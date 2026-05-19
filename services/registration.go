package services

import (
	"context"
	"fmt"
	"reflect"
	"sync"
)

func AddTransient(f any) error {
	return addService(Transient, f)
}

func AddScoped(f any) error {
	return addService(Scoped, f)
}

func AddSingleton(f any) error {
	fVal := reflect.ValueOf(f)
	if fVal.Kind() == reflect.Func && fVal.Type().NumOut() == 1 {
		var results []reflect.Value
		once := sync.Once{}
		wrapper := reflect.MakeFunc(
			fVal.Type(),
			func(args []reflect.Value) []reflect.Value {
				once.Do(func() {
					results = invokeFunction(context.Background(), fVal)
				})
				return results
			})
		return addService(Singleton, wrapper.Interface())
	}
	return fmt.Errorf("function can't be singleton: %v", fVal.Type().Name())
}
