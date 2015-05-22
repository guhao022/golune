package models

import "gopkg.in/mgo.v2/bson"

type (
	Admin struct {
		Id       bson.ObjectId `bson:"_id,omitempty"`
		Username string        `bson:"username,omitempty" valid:"Required;MinSize(2);MaxSize(22)"`
		Password string        `bson:"password" valid:"Required;Range(6, 30)"`
	}
	Blog struct {
		Id              bson.ObjectId `bson:"_id,omitempty"`
		Htmltag         string        `bson:"html"`
		Title           string        `bson:"title" valid:"Required;MinSize(2);MaxSize(22)"`
		Content         string        `bson:"content" valid:"Required"`
		Author          string        `bson:"author"`
		Status          int           `bson:"status"`
		Category        string        `bson:"cate"`
		Hot             int           `bson:"hot"`
		Ip              string        `bson:"ip"`
		MateTitle       string        `bson:"mate_title`
		MateKeyword     string        `bson:"mate_keyword"`
		MateDescription string        `bson:"mate_description"`
		CreatedAt       int64         `bson:"created_at"`
		UpdatedAt       int64         `bson:"updated_at"`
		DeletedAt       int64         `bson:"deleted_at"`
		TagBox          []Tag         `bson:"-"`
		CateBox         string        `bson:"-"`
		CreatedYear     int           `bson:"-"`
		CreatedMonth    int           `bson:"-"`
		CreatedDay      int           `bson:"-"`
	}
	Tag struct {
		Id   bson.ObjectId `bson:"_id,omitempty"`
		Name string        `bson:"name"`
		Num  int           `bson:"num"`
	}
	BlogTag struct {
		Id     bson.ObjectId `bson:"_id,omitempty"`
		BlogId bson.ObjectId `bson:"blog"`
		TagId  bson.ObjectId `bson:"tag"`
	}
)
