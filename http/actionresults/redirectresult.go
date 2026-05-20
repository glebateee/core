package actionresults

import "net/http"

type RedirectAction struct {
	url string
}

func (r *RedirectAction) Execute(ctx *ActionContext) error {
	ctx.ResponseWriter.Header().Set("Location", r.url)
	ctx.ResponseWriter.WriteHeader(http.StatusSeeOther)
	return nil
}

func NewRedirectAction(url string) ActionResult {
	return &RedirectAction{url: url}
}
