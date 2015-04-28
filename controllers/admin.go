package controllers

import (
	"fmt"
	"golune/models"
)

type AdminController struct {
	BaseController
}

func (c *AdminController) Add() {
	var a models.Admin
	a.Userame = "guhao022"
	a.Password = Md5WithSalt("378999587", "jdHl8G9ygR72")
	if err := models.FindByName("guhao022"); err != nil {
		if err := models.AddAdmin(&a); err == nil {
			c.Data["json"] = map[string]int{"Recode": 200}
		} else {
			fmt.Println(fmt.Sprintf("%s", err))
			c.Data["json"] = fmt.Sprintln(err)
		}
	} else {
		c.Data["json"] = map[string]int{"Recode": 201}
	}
}
