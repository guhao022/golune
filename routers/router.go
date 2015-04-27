package routers

import (
	"github.com/astaxie/beego"
	"golune/controllers"
)

func init() {
	beego.Router("/admin", &controllers.AdminController{}, "get:Add")
}
