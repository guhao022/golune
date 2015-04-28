package models

import (
	"golune/mgo"
	"gopkg.in/mgo.v2/bson"
)

func AddAdmin(m *Admin) (err error) {
	db := mgo.NewDB("admin")
	err = db.Insert(m)
	return
}

func FindByName(username string) (err error) {
	var m []Admin
	db := mgo.NewDB("admin")
	err = db.Find(bson.M{"username": username}).All(&m)
	return
}
