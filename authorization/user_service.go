package authorization

import (
	"github.com/glebateee/core/authorization/identity"
	"github.com/glebateee/core/services"
	"github.com/glebateee/core/sessions"
)

func RegisterDefaultUserService() {
	if err := services.AddScoped(func(session sessions.Session, store identity.UserStore) identity.User {
		id, ok := session.GetValue(USER_SESSION_KEY).(int)
		if !ok {
			return identity.UnauthenticatedUser
		}
		user, found := store.GetUserByID(id)
		if !found {
			return identity.UnauthenticatedUser
		}
		return user
	}); err != nil {
		panic(err)
	}
}
