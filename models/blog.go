package models

import (
	"github.com/astaxie/beego/validation"
	"gopkg.in/mgo.v2/bson"
)

func (blog *Blog) Validate(v *validation.Validation) {
	var mgo *Mgo
	if _, err := mgo.FindBlogByName(blog.Title); err == nil {
		v.SetError("Title", "标题重复")
	}
	v.Required(blog.Title, "title").Message("文章标题不能为空")
	v.Range(blog.Title, 2, 50, "title").Message("文章标题在2-50个字符之间")
	v.Required(blog.Category, "cate").Message("分类不能为空")
}

//新建博客
func (mgo *Mgo) NewBlog(m *Blog) (err error) {
	c := mgo.session.DB(BlogDB).C(BlogC)
	err = c.Insert(m)
	return
}

//根据博客名查找博客
func (mgo *Mgo) FindBlogByName(title string) (m *Blog, err error) {
	c := mgo.session.DB(BlogDB).C(AdminC)
	err = c.Find(bson.M{"title": title}).One(&m)
	return m, err
}
