package mgo

import (
	"golune/conf"
	"gopkg.in/mgo.v2"
)

type MgoDB struct {
	MgoHost string
	MgoName string
}

func RegisterDB(url string, dbname string) *mgo.Database {
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	db := session.DB(dbname)
	return db
}

func NewDB(collection string) *mgo.Collection {
	db := RegisterDB(conf.MgoHost, conf.MgoName)
	return db.C(collection)
}
