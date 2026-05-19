package pipeline

import (
	"net/http"
	"reflect"

	"github.com/glebateee/core/services"
)

type RequestPipeline func(*ComponentContext)

var emptyPipeline RequestPipeline = func(cc *ComponentContext) {}

func CreatePipeline(mw ...any) RequestPipeline {
	f := emptyPipeline
	for i := len(mw) - 1; i >= 0; i-- {
		curr := mw[i]
		services.Populate(curr)
		next := f

		if mc, ok := curr.(Middleware); ok {
			f = func(ctx *ComponentContext) {
				if ctx.error == nil {
					mc.ProcessRequest(ctx, next)
				}
			}
			mc.Init()
		} else if smc, ok := curr.(ServicesMiddlewareComponent); ok {
			f = createServiceDependentFunction(curr, next)
			smc.Init()
		}
	}
	return f
}

func (p RequestPipeline) ProcessRequest(r *http.Request, w http.ResponseWriter) error {
	ctx := ComponentContext{Request: r, ResponseWriter: w}
	p(&ctx)
	return ctx.error
}

func createServiceDependentFunction(comp any, next func(*ComponentContext)) RequestPipeline {
	method := reflect.ValueOf(comp).MethodByName("ProcessRequestWithServices")
	if !method.IsValid() {
		panic("ProcessRequestWithServices not implemented")
	}
	return func(ctx *ComponentContext) {
		if ctx.error == nil {
			if _, err := services.CallForContext(
				ctx.Request.Context(),
				method.Interface(),
				ctx,
				next,
			); err != nil {
				ctx.Error(err)
			}
		}
	}
}
