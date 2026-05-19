package services

import (
	"context"
	"fmt"
	"reflect"
)

func GetService(ptr any) error {
	return GetServiceForContext(context.Background(), ptr)
}

func GetServiceForContext(ctx context.Context, ptr any) error {
	ptrVal := reflect.ValueOf(ptr)
	if ptrVal.Kind() == reflect.Pointer && ptrVal.Elem().CanSet() {
		return resolveServiceFromValue(ctx, ptrVal)
	}
	return fmt.Errorf("type is not a pointer: %v", ptrVal.Type().Name())
}
