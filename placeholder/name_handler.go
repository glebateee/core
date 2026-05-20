package placeholder

import (
	"fmt"

	"github.com/glebateee/core/http/actionresults"
	"github.com/glebateee/core/http/handling"
	"github.com/glebateee/core/logging"
	"github.com/glebateee/core/validation"
)

var names = []string{"Alice", "Bob", "Charlie", "Dora"}

type NameHandler struct {
	logging.Logger
	handling.URLGenerator
	validation.Validator
}

func (n NameHandler) GetName(i int) actionresults.ActionResult {
	n.Logger.Debugf("GetName method invoked with argument: %v", i)
	r := "Index out of bounds"
	if i < len(names) {
		r = fmt.Sprintf("Name #%v: %v", i, names[i])
	}
	return actionresults.NewTemplateAction("simple_message.html", r)
}

func (n NameHandler) GetNames() actionresults.ActionResult {
	n.Logger.Debug("GetNames method invoked")
	return actionresults.NewTemplateAction("simple_message.html", fmt.Sprintf("Names: %v", names))
}

func (n NameHandler) GetForm() actionresults.ActionResult {
	postURl, _ := n.URLGenerator.GenerateURL(NameHandler.PostName)
	return actionresults.NewTemplateAction("name_form.html", postURl)
}

type NewName struct {
	Name          string `validation:"required,min:3"`
	InsertAtStart bool
}

func (n NameHandler) PostName(new NewName) actionresults.ActionResult {
	n.Logger.Debugf("PostName method invoked with argument %v", new)
	if ok, errs := n.Validator.Validate(&new); !ok {
		return actionresults.NewTemplateAction("validation_errors.html", errs)
	}
	if new.InsertAtStart {
		names = append([]string{new.Name}, names...)
	} else {
		names = append(names, new.Name)
	}
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetRedirect() actionresults.ActionResult {
	return n.redirectOrError(NameHandler.GetNames)
}

func (n NameHandler) GetJSONData() actionresults.ActionResult {
	return actionresults.NewJSONResult(names)
}

func (n NameHandler) redirectOrError(handler any, args ...any) actionresults.ActionResult {
	url, err := n.URLGenerator.GenerateURL(handler, args...)
	if err != nil {
		return actionresults.NewErrorAction(err)
	}
	return actionresults.NewRedirectAction(url)
}
