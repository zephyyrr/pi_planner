package model

type Id uint

type User struct {
	Id
	Name     string
	Projects []Project
}

func (u User) String() string {
	return u.Name
}
