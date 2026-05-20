package handling

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/http/handling/params"
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/services"
)

type RouterComponent struct {
	routes []Route
}

func NewRouter(handlers ...HandlerEntry) *RouterComponent {
	routes := generateRoutes(handlers...)
	var urlGen URLGenerator
	services.GetService(&urlGen)
	if urlGen == nil {
		services.AddSingleton(func() URLGenerator {
			return &routeURLGenerator{routes: routes}
		})
	} else {
		urlGen.AddRoutes(routes)
	}
	return &RouterComponent{routes: routes}
}

func (router *RouterComponent) Init() {}
func (router *RouterComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	for _, route := range router.routes {
		if strings.EqualFold(ctx.Request.Method, route.httpMethod) {
			matches := route.expression.FindAllStringSubmatch(ctx.URL.Path, -1)
			if len(matches) > 0 {
				rawParamVals := []string{}
				if len(matches[0]) > 1 {
					rawParamVals = matches[0][1:]
				}
				if err := router.invokeHandler(route, rawParamVals, ctx); err != nil {
					ctx.Error(err)

				} else {
					next(ctx)
				}
				return
			}
		}
	}
	ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
}

func (router *RouterComponent) invokeHandler(
	route Route,
	rawParams []string,
	ctx *pipeline.ComponentContext,
) error {
	paramVals, err := params.GetParametersFromRequest(ctx.Request, route.handlerMethod, rawParams)
	if err != nil {
		return err
	}
	structVal := reflect.New(route.handlerMethod.Type.In(0))
	services.PopulateForContext(ctx.Context(), structVal.Interface())
	// TODO
	paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)
	fmt.Println(paramVals[0].Type().Name(), len(paramVals), route.handlerMethod.Name)
	result := route.handlerMethod.Func.Call(paramVals)
	if len(result) > 0 {
		if action, ok := result[0].Interface().(actionresults.ActionResult); ok {
			invoker := createInvokeHandlerFunc(ctx.Context(), router.routes)
			if err := services.PopulateForContextWithExtras(
				ctx.Context(),
				action,
				map[reflect.Type]reflect.Value{
					reflect.TypeOf(invoker): reflect.ValueOf(invoker),
				},
			); err != nil {
				io.WriteString(ctx.ResponseWriter, fmt.Sprint(result[0].Interface()))
				return err
			} else {
				if err := action.Execute(&actionresults.ActionContext{
					Context:        ctx.Context(),
					ResponseWriter: ctx.ResponseWriter,
				}); err != nil {
					return err
				}
			}
		} else {
			io.WriteString(ctx.ResponseWriter, fmt.Sprint(result[0].Interface()))
		}
	}
	return err
}
