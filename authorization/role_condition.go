package authorization

import (
	"slices"

	"github.com/glebateee/core/authorization/identity"
)

type roleCondition struct {
	allowedRoles []string
}

func NewRoleCondition(roles ...string) identity.AuthorizationCondition {
	return &roleCondition{allowedRoles: roles}
}

func (c *roleCondition) Validate(user identity.User) bool {
	return slices.ContainsFunc(c.allowedRoles, user.HasRole)
}
