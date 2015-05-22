package apis

import (
	"fmt"
	"golune/models"
)

type ReadHandler struct {
	BaseHandler
}

//用户登录
func (c *ReadHandler) Login() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	username := c.GetString("username")
	password := c.GetString("password")
	if err := mgo.CheckAdmin(username, password); err == nil {
		c.Data["json"] = map[string]int{"Recode": 200}
	} else {
		c.CustomAbort(403, "login faild")
	}
	c.ServeJson()
}

//获取所有blog
func (c *ReadHandler) BlogList() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	if blogs, err := mgo.FindAllBlog(); err == nil {
		c.Data["json"] = blogs
	} else {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	c.ServeJson()
}

//根据cate获取所有分类下的文章
/*
func (c *ReadHandler)GetBlogByCate(){
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	username := c.GetString("username")
	if blogs, err := mgo.FindBlogByCate(); err == nil {
		c.Data["json"] = blogs
	} else {
		c.CustomAbort(403, fmt.Sprintln(err))
	}
	c.ServeJson()
}*/
