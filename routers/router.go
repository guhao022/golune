package routers

import (
	"github.com/astaxie/beego"
	"golune/controllers"
)

func init() {
	beego.Include(&controllers.AdminController{})
	//beego.Router("/admin/add", &controllers.AdminController{}, "get:Add")
}
