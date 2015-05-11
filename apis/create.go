package apis

import (
	"fmt"
	"golune/models"
)

type CreateHandler struct {
	BaseHandler
}

func (c *CreateHandler) CreateAdmin() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, "/Err500")
	}
	defer mgo.Close()
	var a models.Admin
	a.Username = "guhao022"
	a.Password = Md5WithSalt("378999587", "jdHl8G9ygR72")
	if _, err := mgo.FindAdminByName("guhao022"); err != nil {
		if err := mgo.AddAdmin(&a); err == nil {
			c.Data["json"] = map[string]string{"Recode": "200"}
		} else {
			fmt.Println(fmt.Sprintf("%s", err))
			c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
		}
	} else {
		c.Data["json"] = map[string]string{"Recode": "201"}
	}
	c.ServeJson()
}

func (c *CreateHandler) CreateBlog() {
	mgo, err := models.NewDB()
	if err != nil {
		c.Ctx.Redirect(500, "/Err500")
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
	if _, err := mgo.FindBlogByName(title); err == nil {
		c.Data["json"] = map[string]string{"ErrorMsg": "已存在博客名！"}
	}
}
