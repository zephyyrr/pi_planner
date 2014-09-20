package model

type Project struct {
	Members    map[Id]Rights
	Activities []Activity
	Resources  []Resource
}

type Rights struct {
	Read, Write, Admin bool
}
