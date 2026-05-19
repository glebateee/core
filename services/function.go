package services

import (
	"context"
	"fmt"
	"reflect"
)

func Call(tgt any, args ...any) ([]any, error) {
	return CallForContext(context.Background(), tgt, args...)
}
func CallForContext(ctx context.Context, tgt any, args ...any) ([]any, error) {
	tgtVal := reflect.ValueOf(tgt)
	if tgtVal.Kind() != reflect.Func {
		fmt.Printf("------------------------%s not a func\n", tgtVal.Type().Name())
		return nil, fmt.Errorf("not a function: %v", tgtVal.Type().Name())
	}
	callRes := invokeFunction(ctx, tgtVal, args...)
	callResAny := make([]any, len(callRes))
	for i := range callResAny {
		callResAny[i] = callRes[i].Interface()
	}
	return callResAny, nil
}
