package model

type Model interface {
	Organizations() []Organization
	Organization(id string) Organization

	Users() []User
	User(id string) User
}
