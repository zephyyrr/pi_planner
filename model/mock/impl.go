package mock

import (
	. "github.com/zephyyrr/pplanneer/model"
)

type MockModel struct{}

func Instance() Model {
	return MockModel{}
}

func (m MockModel) Users() []User {
	return []User{m.User("zephyyrr")}
}

func (MockModel) User(id string) User {
	return User{
		Name: id,
	}
}

func (m MockModel) Organizations() []Organization {
	return []Organization{m.Organization("BreamIO")}
}

func (m MockModel) Organization(id string) Organization {
	return Organization{Name: id,
		Members: []User{
			m.User("zephyyrr"),
		},
		Projects: []Project{},
	}
}
