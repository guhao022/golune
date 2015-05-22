package apis

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"golune/models"
)

type UpdateHandler struct {
	BaseHandler
}

//更新博客访问量
func (c *UpdateHandler) BlogHot() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()

	blogId := c.GetString("blogid")
	err = mgo.UpdateBlogAccess(blogId)
	if err != nil {
		c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
	} else {
		c.Data["json"] = map[string]string{"Recode": "200"}
	}
	c.ServeJson()
}

//更新博客
func (c *UpdateHandler) UpdateBlog() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, fmt.Sprintln(err))
	}
	defer mgo.Close()

	var blog models.Blog
	blogid := c.GetString("blogid")
	title := c.GetString("title")
	category := c.GetString("category")
	tags := c.GetString("tags")
	mateTitle := c.GetString("mate_title")
	mateKeywords := c.GetString("mate_keywords")
	mateDescription := c.GetString("mate_description")
	content := c.GetString("content")
	b, err := mgo.FindBlogByName(title)
	if err == nil && blogid != b.Id.Hex() {
		c.Data["json"] = map[string]string{"Recode": "400", "Error": "已存在博客名"}
	} else {
		blog.Title = title
		blog.Category = category
		blog.MateTitle = mateTitle
		blog.MateKeyword = mateKeywords
		blog.MateDescription = mateDescription
		blog.Content = content
		//检测字段是否符合要求 validate
		valid := validation.Validation{}
		b, err := valid.Valid(&blog)
		if err != nil {
			c.CustomAbort(500, fmt.Sprintln(err))
		}
		if !b {
			c.CustomAbort(403, fmt.Sprintln(valid.ErrorsMap))
		}
		//写入数据库
		if err := mgo.NewBlog(&blog); err == nil {
			mgo.DelBlogTagByBlog(blogid)
			//循环tags并写入tags表
			c.CreateTags(tags, blog.Id)
			c.Data["json"] = map[string]int{"Recode": 200}
		} else {
			c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
		}
	}
	c.ServeJson()
}
