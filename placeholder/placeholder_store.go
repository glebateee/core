package placeholder

import (
	"strings"

	"github.com/glebateee/core/authorization/identity"
	"github.com/glebateee/core/services"
)

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
	2: identity.NewBasicUser(2, "Bob"),
}

type PlaceholderUserStore struct{}

func (p *PlaceholderUserStore) GetUserByID(id int) (identity.User, bool) {
	user, found := users[id]
	return user, found
}

func (p *PlaceholderUserStore) GetUserByName(name string) (user identity.User, found bool) {
	for _, user := range users {
		if strings.EqualFold(user.GetName(), name) {
			return user, true
		}
	}
	return nil, false
}

func RegisterPlaceholderUserStore() {
	err := services.AddSingleton(func() identity.UserStore {
		return &PlaceholderUserStore{}
	})
	if err != nil {
		panic(err)
	}
}
