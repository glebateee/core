package basic

import (
	"net/http"
	"strings"

	"github.com/glebateee/core/config"
	"github.com/glebateee/core/pipeline"
)

type StaticFileComponent struct {
	prefix  string
	handler http.Handler
	Cfg     config.Config
}

func (c *StaticFileComponent) Init() {
	c.prefix = c.Cfg.GetStringDefault("files:urlprefix", "/files/")
	path, ok := c.Cfg.GetString("files:path")
	if !ok {
		panic("static file handler configuration not found")
	}
	c.handler = http.StripPrefix(c.prefix, http.FileServer(http.Dir(path)))
}

func (c *StaticFileComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	url := ctx.Request.URL.Path
	if !strings.EqualFold(url, c.prefix) && strings.HasPrefix(strings.ToLower(url), strings.ToLower(c.prefix)) {
		c.handler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	} else {
		next(ctx)
	}
}
