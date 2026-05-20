package templates

import (
	"html/template"
	"io"
	"strings"
)

type LayoutTemplateProcessor struct{}

var getTemplates func() *template.Template

func insertBodyWrapper(b *strings.Builder) func() template.HTML {
	return func() template.HTML {
		return template.HTML(b.String())
	}
}

func setLayoutWrapper(layout *string) func(string) string {
	return func(name string) string {
		*layout = name
		return ""
	}
}

var emptyFunc = func(handlerName, methodName string, args ...any) any { return "" }

func (p *LayoutTemplateProcessor) ExecTemplate(w io.Writer, name string, data any) error {
	return p.ExecTemplateWithFunc(w, name, data, emptyFunc)
}
func (p *LayoutTemplateProcessor) ExecTemplateWithFunc(w io.Writer, name string, data any, f InvokeHandlerFunc) error {
	var layoutName string
	var sb strings.Builder
	localTemplates := getTemplates()
	localTemplates.Funcs(template.FuncMap{
		"body":    insertBodyWrapper(&sb),
		"layout":  setLayoutWrapper(&layoutName),
		"handler": f,
	})

	if err := localTemplates.ExecuteTemplate(&sb, name, data); err != nil {
		return err
	}
	if layoutName == "" {
		_, err := io.WriteString(w, sb.String())
		if err != nil {
			return err
		}
	} else {
		if err := localTemplates.ExecuteTemplate(w, layoutName, data); err != nil {
			return err
		}
	}
	return nil
}
