package routers

import (
	"github.com/astaxie/beego"
	"golune/apis"
)

func init() {
	ns :=
		beego.NewNamespace("/api",
			beego.NSNamespace("/blog",
				//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
				beego.NSNamespace("/create",
					beego.NSRouter("/admin", &apis.CreateHandler{}, "get:CreateAdmin"),
					//beego.NSRouter("/login", &apis.AdminHandler{}, "post:Login"),
				),
				beego.NSNamespace("/read",
					beego.NSRouter("/login", &apis.ReadHandler{}, "get:Login"),
					//beego.NSRouter("/login", &apis.AdminHandler{}, "post:Login"),
				),
			),
		)

	beego.AddNamespace(ns)
	//beego.Router("/api", &apis.AdminHandler{}, "get,post:AddAdmin")
}
