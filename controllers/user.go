package controllers

import (
	"github.com/astaxie/beego"
	"golune/models"
)

type AdminController struct {
	beego.Controller
}

func (c AdminController) AddAdmin() {
	var a models.Admin
}
