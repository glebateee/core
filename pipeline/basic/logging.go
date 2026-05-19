package basic

import (
	"net/http"

	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/pipeline"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *LoggingResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *LoggingResponseWriter) Write(b []byte) (int, error) {
	if w.statusCode == 0 {
		w.statusCode = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

type LoggingComponent struct{}

func (c *LoggingComponent) Init() {}

func (c *LoggingComponent) ImplementsProcessRequestWithServices() {}

func (c *LoggingComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	logger logging.Logger,
) {
	lw := LoggingResponseWriter{ResponseWriter: ctx.ResponseWriter}
	ctx.ResponseWriter = &lw
	// for name, values := range ctx.Request.Header {
	// 	fmt.Printf("Header: %s, Value: %s\n", name, strings.Join(values, ", "))
	// }
	logger.Infof("REQ  ---  %v  -  %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	// for name, values := range ctx.ResponseWriter.Header() {
	// 	fmt.Printf("Header: %s, Value: %s\n", name, strings.Join(values, ", "))
	// }
	logger.Infof("RSP  ---  %v  -  %v", lw.statusCode, ctx.Request.URL)

}
