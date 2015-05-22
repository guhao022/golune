package models

import (
	//"github.com/russross/blackfriday"
	"golune/utils"
	"gopkg.in/mgo.v2/bson"

	"time"
)

//新建博客
func (mgo *Mgo) NewBlog(m *Blog) (err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	t := time.Now()
	m.Author = "Halgu"
	m.Hot = 0
	//1为显示
	m.Status = 1
	m.CreatedAt = t.Unix()
	m.CreatedYear = t.Year()
	m.CreatedMonth = utils.Month2Int(t.Month())
	m.CreatedDay = t.Day()
	err = c.Insert(m)
	return
}

//根据博客名查找博客
func (mgo *Mgo) FindBlogByName(title string) (m *Blog, err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	err = c.Find(bson.M{"title": title}).One(&m)
	if err == nil {
		return m, nil
	} else {
		return nil, err
	}
}

//根据cate查找博客
func (mgo *Mgo) FindBlogByCate(cate string) (m []Blog, err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	query := c.Find(bson.M{"cate": cate}).Sort("-created_at")
	err = query.All(&m)
	return
}

//根据tag查找blog
func (mgo *Mgo) FindBlogTagByTags(tagId bson.ObjectId) (m []BlogTag, err error) {
	c := mgo.session.DB(BlogDB).C(BlogTagC)
	query := c.Find(bson.M{"tag": tagId}).Sort("-created_at")
	err = query.All(&m)
	return
}

//根据ID查找blog
func (mgo *Mgo) FindBlogById(id string) (m []Blog, err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	query := c.FindId(bson.M{"_id": id})
	err = query.All(&m)
	return
}

//根据发表时间（年）查找blog
func (mgo *Mgo) FindBlogByYear(year int) (m []Blog, err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	query := c.Find(bson.M{"create_year": year}).Sort("-created_at")
	err = query.All(&m)
	return
}

//查找所有已发表博客
func (mgo *Mgo) FindAllBlog() (m []Blog, err error) {
	//var m []Blog
	c := mgo.session.DB(BlogDB).C(BlogC)
	query := c.Find(bson.M{"status": 1, "deleted_at": 0}).Sort("-created_at")
	err = query.All(&m)
	return
}
