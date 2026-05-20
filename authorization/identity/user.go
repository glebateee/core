package identity

type User interface {
	GetID() int
	GetName() string
	HasRole(name string) bool
	IsAuthenticated() bool
}
