package handling

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strings"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/services"
	"github.com/glebateee/core/templates"
)

func createInvokeHandlerFunc(
	ctx context.Context,
	routes []Route,
) templates.InvokeHandlerFunc {
	return func(handlerName, methodName string, args ...any) any {
		var anyRoute bool
		for _, route := range routes {
			if strings.EqualFold(handlerName, route.handlerName) &&
				strings.EqualFold(methodName, route.handlerMethod.Name) {
				anyRoute = true
				paramVals := make([]reflect.Value, len(args))
				for i := range args {
					paramVals[i] = reflect.ValueOf(args[i])
				}
				structVal := reflect.New(route.handlerMethod.Type.In(0))
				services.PopulateForContext(ctx, structVal.Interface())
				paramVals = append([]reflect.Value{structVal.Elem()}, paramVals...)
				result := route.handlerMethod.Func.Call(paramVals)
				if action, ok := result[0].Interface().(*actionresults.TemplateActionResult); ok {
					invoker := createInvokeHandlerFunc(ctx, routes)
					if err := services.PopulateForContextWithExtras(
						ctx,
						action,
						map[reflect.Type]reflect.Value{
							reflect.TypeOf(invoker): reflect.ValueOf(invoker),
						}); err != nil {
						return fmt.Errorf("PopulateForContextWithExtras rendering error: %w", err)
					}
					w := &stringsResponseWriter{Builder: &strings.Builder{}}
					if err := action.Execute(&actionresults.ActionContext{
						Context:        ctx,
						ResponseWriter: w,
					}); err != nil {
						return fmt.Errorf("Execute error: %w", err)
					}
					return (template.HTML)(w.Builder.String())
				} else {
					return fmt.Sprint(result[0])
				}
			}
		}
		if !anyRoute {
			return fmt.Errorf("no route found")
		}
		return nil
	}
}

type stringsResponseWriter struct {
	*strings.Builder
}

func (s *stringsResponseWriter) Write(b []byte) (int, error) {
	return s.Builder.Write(b)
}

func (s *stringsResponseWriter) WriteHeader(int)     {}
func (s *stringsResponseWriter) Header() http.Header { return http.Header{} }
