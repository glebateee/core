package actionresults

type ErrorAction struct {
	error
}

func (e *ErrorAction) Execute(ctx *ActionContext) error {
	return e.error
}

func NewErrorAction(err error) ActionResult {
	return &ErrorAction{error: err}
}
