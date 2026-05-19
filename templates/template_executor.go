package templates

import "io"

type TemplateExecutor interface {
	ExecTemplate(w io.Writer, name string, data ...any) error
}
