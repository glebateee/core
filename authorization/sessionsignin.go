package authorization

import (
	"context"

	"github.com/glebateee/core/authorization/identity"
	"github.com/glebateee/core/services"
	"github.com/glebateee/core/sessions"
)

type SessionSignInMgr struct {
	context.Context
}

const USER_SESSION_KEY string = "USER"

func (m *SessionSignInMgr) SignIn(user identity.User) error {
	session, err := m.getSession()
	if err != nil {
		return err
	}
	session.SetValue(USER_SESSION_KEY, user.GetID())
	return nil
}

func (m *SessionSignInMgr) SignOut(user identity.User) error {
	session, err := m.getSession()
	if err != nil {
		return err
	}
	session.SetValue(USER_SESSION_KEY, nil)
	return nil
}

func (m *SessionSignInMgr) getSession() (sessions.Session, error) {
	var s sessions.Session
	if err := services.GetServiceForContext(m.Context, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func RegisterDefaultSignInService() {
	if err := services.AddScoped(func(ctx context.Context) identity.SignInManager {
		return &SessionSignInMgr{Context: ctx}
	}); err != nil {
		panic(err)
	}
}
