package routers

import (
	"github.com/astaxie/beego"
	"golune/controllers"
)

func init() {
	beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
}
