package mongo

import (
	. "github.com/zephyyrr/pplanneer/model"

	mgo "gopkg.in/mgo.v2"
)

type MongoModel struct {
	session *mgo.Session
}

func (mm MongoModel) Users() []User {
	mm.session.
}
