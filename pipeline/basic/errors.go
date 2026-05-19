package basic

import (
	"net/http"

	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/pipeline"
)

type ErrorComponent struct{}

func (c *ErrorComponent) Init() {}

func (c *ErrorComponent) ImplementsProcessRequestWithServices() {}

func (c *ErrorComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {

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
