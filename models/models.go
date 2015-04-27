package models

import "gopkg.in/mgo.v2/bson"

type Admin struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Userame  string        `json:"username"`
	Password int           `json:"password"`
}

type Blog struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Title           string        `json:"title"`
	Content         string        `json:"title"`
	AuthorId        bson.ObjectId `json:"author"`
	Status          int           `json:"status"`
	CateId          bson.ObjectId `json:"cate"`
	Hot             int           `json:"hot"`
	Ip              string        `json:"ip"`
	MateTitle       string        `json:"mate_title`
	MateKeyword     string        `json:"mate_keyword"`
	MateDescription string        `json:"mate_description"`
	CreatedAt       int           `json:"created_at"`
	UpdatedAt       int           `json:"updated_at"`
	DeletedAt       int           `json:"deleted_at"`
}

type Tag struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
	Num  int           `json:"num"`
}

type BlogTag struct {
	Id     bson.ObjectId `json:"id" bson:"_id"`
	BlogId bson.ObjectId `json:"blog"`
	TagId  bson.ObjectId `json:"tag"`
}
