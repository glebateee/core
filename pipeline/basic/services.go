package basic

import (
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/services"
)

type ServicesComponent struct{}

func (c *ServicesComponent) Init() {}

func (c *ServicesComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	rc := ctx.Request.Context()
	ctx.Request = ctx.Request.WithContext(services.NewServiceContext(rc))
	next(ctx)
}
