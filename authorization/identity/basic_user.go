package identity

import "strings"

var UnauthenticatedUser User = &basicUser{}

func NewBasicUser(id int, name string, roles ...string) User {
	return &basicUser{
		ID:            id,
		Name:          name,
		Roles:         roles,
		Authenticated: true,
	}
}

type basicUser struct {
	ID            int
	Name          string
	Roles         []string
	Authenticated bool
}

func (user *basicUser) GetID() int {
	return user.ID
}
func (user *basicUser) GetName() string {
	return user.Name
}
func (user *basicUser) HasRole(role string) bool {
	for _, r := range user.Roles {
		if strings.EqualFold(r, role) {
			return true
		}
	}
	return false
}

func (user *basicUser) IsAuthenticated() bool {
	return user.Authenticated
}
