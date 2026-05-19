package placeholder

import (
	"github.com/glebateee/core/config"
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/templates"
)

type SimpleMessageComponent struct {
	msg string
	config.Config
}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (c *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	exec templates.TemplateExecutor,
) {
	if err := exec.ExecTemplate(ctx.ResponseWriter, "simple_message.html", c.msg); err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
