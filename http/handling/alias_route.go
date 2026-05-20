package handling

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/services"
)

func (c *RouterComponent) AddMethodAlias(
	srcURL string,
	dstMethod any,
	args ...any,
) *RouterComponent {
	var genURL URLGenerator
	if err := services.GetService(&genURL); err != nil {
		panic(err)
	}
	dstURL, err := genURL.GenerateURL(dstMethod, args...)
	if err != nil {
		panic(err)
	}
	return c.AddURLAlias(srcURL, dstURL)
}

func (c *RouterComponent) AddURLAlias(
	srcURL string,
	tgtURL string,
) *RouterComponent {
	calledMethod := func(any) actionresults.ActionResult {
		return actionresults.NewRedirectAction(tgtURL)
	}
	alias := Route{
		httpMethod: http.MethodGet,
		//handlerName: "Alias",
		//actionName:  "Redirect",
		expression: *regexp.MustCompile(fmt.Sprintf("^%v[/]?$", srcURL)),
		handlerMethod: reflect.Method{
			Type: reflect.TypeOf(calledMethod),
			Func: reflect.ValueOf(calledMethod),
		},
	}
	c.routes = append([]Route{alias}, c.routes...)
	return c
}
