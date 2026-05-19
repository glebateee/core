package basic

import (
	"fmt"
	"net/http"
	"strings"

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
	// var logger logging.Logger
	// if err := services.GetServiceForContext(ctx.Request.Context(), &logger); err != nil {
	// 	ctx.Error(err)
	// 	return
	// }
	lw := LoggingResponseWriter{ResponseWriter: ctx.ResponseWriter}
	ctx.ResponseWriter = &lw
	//fmt.Println(ctx.Request.Header.Get("ETag"), ctx.Request.Header.Get("Last-Modified"), ctx.Request.Header.Get("Cache-Control"))
	for name, values := range ctx.Request.Header {
		// Имена заголовков (name) регистронезависимы, но для вывода можно привести к стандартному виду
		fmt.Printf("Header: %s, Value: %s\n", name, strings.Join(values, ", "))
	}
	logger.Infof("REQ  ---  %v  -  %v", ctx.Request.Method, ctx.Request.URL)
	next(ctx)
	for name, values := range ctx.ResponseWriter.Header() {
		// Имена заголовков (name) регистронезависимы, но для вывода можно привести к стандартному виду
		fmt.Printf("Header: %s, Value: %s\n", name, strings.Join(values, ", "))
	}
	//fmt.Println(lw.Header().Get("Etag"), lw.Header().Get("Last-Modified"), lw.Header().Get("Cache-Control"))
	logger.Infof("RSP  ---  %v  -  %v", lw.statusCode, ctx.Request.URL)

}
