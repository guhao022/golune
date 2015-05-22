package models

import (
	"gopkg.in/mgo.v2/bson"
)

//添加标签
func (mgo *Mgo) AddTag(m *Tag) (err error) {
	c := mgo.session.DB(BlogDB).C(TagC)
	err = c.Insert(m)
	return
}

//获取所有tag
func (mgo *Mgo) GetTags() (m []Tag, err error) {
	c := mgo.session.DB(BlogDB).C(TagC)
	err = c.Find(nil).All(&m)
	return
}

//根据标签名查找标签
func (mgo *Mgo) FindTagByName(tag string) (m *Tag, err error) {
	c := mgo.session.DB(BlogDB).C(TagC)
	err = c.Find(bson.M{"name": tag}).One(&m)
	return m, err
}

//添加blogtag
func (mgo *Mgo) AddBlogTag(m *BlogTag) (err error) {
	c := mgo.session.DB(BlogDB).C(BlogTagC)
	err = c.Insert(m)
	return
}

//根据blog查找blogtag
func (mgo *Mgo) FindBlogTagByBlog(blogId bson.ObjectId) (tag []Tag, err error) {
	var m []BlogTag
	var t Tag
	c := mgo.session.DB(BlogDB).C(BlogTagC)
	err = c.Find(bson.M{"blog": blogId}).All(&m)
	for _, blogtag := range m {
		tc := mgo.session.DB(BlogDB).C(TagC)
		tc.Find(bson.M{"_id": blogtag.TagId}).One(&t)
		tag = append(tag, t)
	}
	return
}

//根据blog删除blogtag
func (mgo *Mgo) DelBlogTagByBlog(blogId bson.ObjectId) (err error) {
	c := mgo.session.DB(BlogDB).C(BlogTagC)
	_, err = c.RemoveAll(bson.M{"blog": blogId})
	return
}

//更新tag表信息
func (mgo *Mgo) UpdateTag(m *Tag) (err error) {
	c := mgo.session.DB(BlogDB).C(TagC)
	err = c.Update(bson.M{"_id": m.Id}, bson.M{"$set": bson.M{"num": m.Num, "name": m.Name}})
	return
}
