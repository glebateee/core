package services

import (
	"context"
	"fmt"
	"reflect"
)

func Populate(tgt any) error {
	return PopulateForContext(context.Background(), tgt)
}

func PopulateForContext(
	ctx context.Context,
	tgt any,
) error {
	return PopulateForContextWithExtras(ctx, tgt, make(map[reflect.Type]reflect.Value))
}

func PopulateForContextWithExtras(
	ctx context.Context,
	tgt any,
	extras map[reflect.Type]reflect.Value,
) error {
	tgtVal := reflect.ValueOf(tgt)
	if tgtVal.Kind() == reflect.Pointer && tgtVal.Elem().Kind() == reflect.Struct {
		tgtVal := tgtVal.Elem()
		for _, field := range tgtVal.Fields() {
			if field.CanSet() {
				if extra, ok := extras[field.Type()]; ok {
					field.Set(extra)
				} else {
					resolveServiceFromValue(ctx, field.Addr())
				}
			}
		}
		return nil
	}
	return fmt.Errorf("type cannot be used as target: %v", tgtVal.Type().Name())
}
