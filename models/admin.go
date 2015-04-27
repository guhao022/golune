package models

import (
	"golune/mgo"
)

func AddAdmin(m *Admin) (err error) {
	db := mgo.NewDB("admin")
	err = db.Insert(m)
	return
}

func FindByName(username string) {
	//
}
