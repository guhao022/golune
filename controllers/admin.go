package controllers

import (
	"fmt"
	"golune/models"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) URLMapping() {
	c.Mapping("AddAdmin", c.AddAdmin)
	c.Mapping("Login", c.Login)
}

// @router /add [post]
func (c *AdminController) AddAdmin() {
	var a models.Admin
	a.Userame = "guhao022"
	a.Password = Md5WithSalt("378999587", "jdHl8G9ygR72")
	if err := models.FindByName("guhao022"); err != nil {
		if err := models.AddAdmin(&a); err == nil {
			c.Data["json"] = map[string]string{"Recode": "200"}
		} else {
			fmt.Println(fmt.Sprintf("%s", err))
			c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
		}
	} else {
		c.Data["json"] = map[string]string{"Recode": "201"}
	}
}

// @router /login [post]
func (c *AdminController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	if err := models.CheckAdmin(username, password); err == nil {
		c.Data["json"] = map[string]string{"Recode": "200"}
	} else {
		c.Data["json"] = map[string]string{"Recode": "400", "Error": fmt.Sprintln(err)}
	}
}
