package actionresults

import "github.com/glebateee/core/templates"

type TemplateActionResult struct {
	tamplateName string
	data         any
	templates.TemplateExecutor
	templates.InvokeHandlerFunc
}

func (t *TemplateActionResult) Execute(ctx *ActionContext) error {
	return t.ExecTemplateWithFunc(ctx.ResponseWriter, t.tamplateName, t.data, t.InvokeHandlerFunc)
}

func NewTemplateAction(
	name string,
	data any,
) ActionResult {
	return &TemplateActionResult{
		tamplateName: name,
		data:         data,
	}
}
