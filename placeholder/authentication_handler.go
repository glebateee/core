package placeholder

import (
	"fmt"

	"github.com/glebateee/core/authorization/identity"
	"github.com/glebateee/core/http/actionresults"
)

type AuthenticationHandler struct {
	identity.User
	identity.SignInManager
	identity.UserStore
}

func (h AuthenticationHandler) GetSignIn() actionresults.ActionResult {
	return actionresults.NewTemplateAction("signin.html", fmt.Sprintf("Signed in as: %v", h.User.GetName()))
}

type Credentials struct {
	Username string
	Password string
}

func (h AuthenticationHandler) PostSignIn(creds Credentials) actionresults.ActionResult {
	if creds.Password == "mysecret" {
		user, ok := h.UserStore.GetUserByName(creds.Username)
		if ok {
			h.SignInManager.SignIn(user)
			return actionresults.NewTemplateAction("signin.html", fmt.Sprintf("Signed in as: %v", user.GetName()))
		}
	}
	return actionresults.NewTemplateAction("signin.html", "Access Denied")
}

func (h AuthenticationHandler) PostSignOut() actionresults.ActionResult {
	h.SignInManager.SignOut(h.User)
	return actionresults.NewTemplateAction("signin.html", "Signed out")
}
