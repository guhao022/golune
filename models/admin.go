package models

import (
	"golune/mgo"
	"gopkg.in/mgo.v2/bson"
)

//添加管理员
func AddAdmin(m *Admin) (err error) {
	db := mgo.NewDB("admin")
	err = db.Insert(m)
	return
}

//根据用户名查找用户信息
func FindByName(username string) (err error) {
	var m []Admin
	db := mgo.NewDB("admin")
	err = db.Find(bson.M{"username": username}).All(&m)
	return
}

//检测用户是否匹配
func CheckAdmin(username, password string) (err error) {
	var m []Admin
	db := mgo.NewDB("admin")
	err = db.Find(bson.M{"username": username, "password": password}).One(&m)
	return
}
