package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(w io.Writer, name string, data any) error
	ExecTemplateWithFunc(w io.Writer, name string, data any, handler InvokeHandlerFunc) error
}

type InvokeHandlerFunc func(handlerName, methodName string, args ...any) any
