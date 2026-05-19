package basic

import (
	"net/http"

	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/pipeline"
	"github.com/glebateee/core/services"
)

type ErrorComponent struct{}

func (c *ErrorComponent) Init() {}

func (c *ErrorComponent) ProcessRequest(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
) {
	var logger logging.Logger
	if err := services.GetServiceForContext(ctx.Context(), &logger); err != nil {
		panic(err)
	}
	defer recoverFunc(ctx, logger)
	next(ctx)
	if err := ctx.GetError(); err != nil {
		logger.Debugf("Error: %w", err)
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}

func recoverFunc(ctx *pipeline.ComponentContext, logger logging.Logger) {
	if arg := recover(); arg != nil {
		logger.Debugf("Error: %v", arg)
		ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
	}
}
