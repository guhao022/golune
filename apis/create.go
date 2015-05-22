package apis

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/validation"
	"golune/models"
	"golune/utils"
	"gopkg.in/mgo.v2/bson"
)

type CreateHandler struct {
	BaseHandler
}

func (c *CreateHandler) CreateAdmin() {
	mgo, err := models.NewDB()
	if err != nil {
		c.CustomAbort(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	var a models.Admin
	a.Username = "guhao022"
	a.Password = Md5WithSalt("378999587", "jdHl8G9ygR72")
	valid := validation.Validation{}
	b, err := valid.Valid(&a)
	if err != nil {
		c.CustomAbort(500, fmt.Sprintln(err))
	}
	if !b {
		c.CustomAbort(403, fmt.Sprintln(valid.ErrorsMap))
	}
	if _, err := mgo.FindAdminByName("guhao022"); err != nil {
		if err := mgo.AddAdmin(&a); err == nil {
			c.Data["json"] = map[string]int{"Recode": 200}
		} else {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
	} else {
		c.CustomAbort(202, "已存在用户名")
	}
	c.ServeJson()
}

func (c *CreateHandler) CreateBlog() {
	mgo, err := models.NewDB()
	if err != nil {
		c.CustomAbort(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	var blog models.Blog
	title := c.GetString("title")
	category := c.GetString("category")
	tags := c.GetString("tags")
	mateTitle := c.GetString("mate_title")
	mateKeywords := c.GetString("mate_keywords")
	mateDescription := c.GetString("mate_description")
	content := c.GetString("content")
	ip := c.GetString("ip")
	if _, err := mgo.FindBlogByName(title); err == nil {
		c.CustomAbort(403, "已存在博客名")
	} else {
		blog.Id = bson.NewObjectId()
		blog.Htmltag = string(utils.RandomCreateBytes(12))
		blog.Title = title
		blog.Category = category
		blog.MateTitle = mateTitle
		blog.MateKeyword = mateKeywords
		blog.MateDescription = mateDescription
		blog.Content = content
		blog.Ip = ip
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
			//循环tags并写入tags表
			c.CreateTags(tags, blog.Id)
			c.Data["json"] = map[string]int{"Recode": 200}
		} else {
			c.CustomAbort(403, fmt.Sprintln(err))
		}
	}
	c.ServeJson()
}

func (c *BaseHandler) CreateTags(tags string, blogId bson.ObjectId) {
	mgo, err := models.NewDB()
	if err != nil {
		c.CustomAbort(500, fmt.Sprintln(err))
	}
	defer mgo.Close()
	var tag models.Tag
	var blog_tag models.BlogTag
	tagArr := strings.SplitN(tags, ",", -1)
	for _, t := range tagArr {
		if hasTag, err := mgo.FindTagByName(t); err != nil {
			tag.Id = bson.NewObjectId()
			tag.Num = 1
			tag.Name = t
			if err := mgo.AddTag(&tag); err != nil {
				c.CustomAbort(500, fmt.Sprintln(err))
			}
			blog_tag.BlogId = blogId
			blog_tag.TagId = tag.Id
			if err := mgo.AddBlogTag(&blog_tag); err != nil {
				c.CustomAbort(500, fmt.Sprintln(err))
			}
		} else {
			hasTag.Num = hasTag.Num + 1
			if err := mgo.UpdateTag(hasTag); err != nil {
				c.CustomAbort(500, fmt.Sprintln(err))
			}
			blog_tag.BlogId = blogId
			blog_tag.TagId = hasTag.Id
			if err := mgo.AddBlogTag(&blog_tag); err != nil {
				c.CustomAbort(500, fmt.Sprintln(err))
			}
		}
	}
}
