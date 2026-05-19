package services

import (
	"context"
	"reflect"
)

const serviceKey = "services"

type ServiceMap map[reflect.Type]reflect.Value

func NewServiceContext(ctx context.Context) context.Context {
	if ctx.Value(serviceKey) == nil {
		return context.WithValue(ctx, serviceKey, make(ServiceMap))
	}
	return ctx
}
