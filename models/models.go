package models

import "gopkg.in/mgo.v2/bson"

type Model struct {
}
type (
	Admin struct {
		Id       bson.ObjectId `bson:"_id,omitempty"`
		Username string        `bson:"username,omitempty"`
		Password string        `bson:"password"`
	}
	Blog struct {
		Id              bson.ObjectId `bson:"_id,omitempty"`
		Title           string        `bson:"title"`
		Content         string        `bson:"content"`
		Author          string        `bson:"author"`
		Status          int           `bson:"status"`
		Category        string        `bson:"cate"`
		Hot             int           `bson:"hot"`
		Ip              string        `bson:"ip"`
		MateTitle       string        `bson:"mate_title`
		MateKeyword     string        `bson:"mate_keyword"`
		MateDescription string        `bson:"mate_description"`
		CreatedAt       int           `bson:"created_at"`
		UpdatedAt       int           `bson:"updated_at"`
		DeletedAt       int           `bson:"deleted_at"`
	}
	Tag struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        `bson:"name,omitempty"`
		Num  int           `bson:"num"`
	}
	BlogTag struct {
		Id     bson.ObjectId `bson:"_id,omitempty"`
		BlogId bson.ObjectId `bson:"blog"`
		TagId  bson.ObjectId `bson:"tag"`
	}
)
