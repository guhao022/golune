package models

import (
	"gopkg.in/mgo.v2/bson"
)

//添加管理员
func (mgo *Mgo) AddAdmin(m *Admin) (err error) {
	c := mgo.session.DB(BlogDB).C(AdminC)
	err = c.Insert(m)
	return
}

//根据用户名查找用户信息
func (mgo *Mgo) FindAdminByName(username string) (m *Admin, err error) {
	c := mgo.session.DB(BlogDB).C(AdminC)
	err = c.Find(bson.M{"username": username}).One(&m)
	return m, err
}

//检测用户是否匹配
func (mgo *Mgo) CheckAdmin(username, password string) (err error) {
	m := Admin{}
	c := mgo.session.DB(BlogDB).C(AdminC)
	err = c.Find(bson.M{"username": username, "password": password}).One(&m)
	return
}
