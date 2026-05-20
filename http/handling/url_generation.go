package handling

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type URLGenerator interface {
	GenerateURL(method any, args ...any) (string, error)
	GenerateURLByName(methodName, handlerName string, args ...any) (string, error)
	AddRoutes(routes []Route)
}

type routeURLGenerator struct {
	routes []Route
}

// ask
func (g *routeURLGenerator) AddRoutes(routes []Route) {
	if g.routes == nil {
		g.routes = routes
	} else {
		g.routes = append(g.routes, routes...)
	}
}

func generateURL(route Route, args ...any) (string, error) {
	url := "/" + route.prefix
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}
	url += strings.ToLower(route.actionName)
	if len(args) > 0 && !strings.EqualFold(route.httpMethod, http.MethodGet) {
		return "", fmt.Errorf("parameters allowed only for GET method")
	}
	if strings.EqualFold(route.httpMethod, http.MethodGet) && len(args)+1 != route.handlerMethod.Type.NumIn() {
		return "", fmt.Errorf("provided args number not enough to call method")
	}
	for _, val := range args {
		url = fmt.Sprintf("%s/%v", url, val)
	}
	return url, nil
}

func (g *routeURLGenerator) GenerateURL(method any, args ...any) (string, error) {
	methodVal := reflect.ValueOf(method)
	if methodVal.Kind() == reflect.Func && methodVal.Type().In(0).Kind() == reflect.Struct {
		for _, route := range g.routes {
			if route.handlerMethod.Func.Pointer() == methodVal.Pointer() {
				return generateURL(route, args...)
			}
		}
	}
	return "", fmt.Errorf("no matching route")
}

func (g *routeURLGenerator) GenerateURLByName(methodName, handlerName string, args ...any) (string, error) {
	for _, route := range g.routes {
		if strings.EqualFold(route.handlerName, handlerName) && strings.EqualFold(route.httpMethod+route.actionName, methodName) {
			return generateURL(route, args...)
		}
	}

	return "", fmt.Errorf("no matching route")
}
