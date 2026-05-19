package placeholder

import (
	"errors"
	"io"

	"github.com/glebateee/core/config"
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/services"
)

type SimpleMessageComponent struct{}

func (c *SimpleMessageComponent) Init() {}

func (c *SimpleMessageComponent) ProcessRequest(ctx *pipeline.ComponentContext, next func(*pipeline.ComponentContext)) {
	var cfg config.Config
	services.GetService(&cfg)
	msg, ok := cfg.GetString("main:message")
	if ok {
		io.WriteString(ctx.ResponseWriter, msg)
	} else {
		ctx.Error(errors.New("config setting for message component not found"))
	}
	next(ctx)
}
