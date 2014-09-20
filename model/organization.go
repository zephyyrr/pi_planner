package model

type Organization struct {
	Name      string
	Projects  []Project
	Members   []User
	Teams     []Team
	Resources []Resource
}

func (o Organization) String() string {
	return o.Name
}

func Organizations() []Organization {
	return []Organization{
		Organization{
			Name:      "Bream iO AB",
			Members:   []User{User{Name: "Zephyyrr"}, User{Name: "Nevonen"}, User{Name: "ThompZon"}},
			Resources: []Resource{Resource{3, "Tobii REX Eye Tracker"}},
		},
	}
}
