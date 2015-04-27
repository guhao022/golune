package models

import (
	"golune/utils/mgo"
	"gopkg.in/mgo.v2/bson"
)

func AddAdmin(m *Admin) (err error) {
	db := mgo.NewDB("admin")
	err = db.Insert(m)
	return
}
