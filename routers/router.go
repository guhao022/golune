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
					beego.NSRouter("/blog", &apis.CreateHandler{}, "post:CreateBlog"),
				),
				beego.NSNamespace("/read",
					beego.NSRouter("/login", &apis.ReadHandler{}, "get:Login"),
					beego.NSRouter("/blog", &apis.ReadHandler{}, "get:BlogList"),
				),
				beego.NSNamespace("/cache",
					beego.NSRouter("/blog", &apis.OutputCacheHandler{}, "get:CacheBlogHtml"),
					beego.NSRouter("/index", &apis.OutputCacheHandler{}, "get:CacheIndexHtml"),
					beego.NSRouter("/cate", &apis.OutputCacheHandler{}, "get:CacheCateHtml"),
					beego.NSRouter("/tags", &apis.OutputCacheHandler{}, "get:CacheTagsHtml"),
					beego.NSRouter("/archive", &apis.OutputCacheHandler{}, "get:CacheArchiveHtml"),
				),
			),
		)

	beego.AddNamespace(ns)
	//beego.Router("/api", &apis.OutputCacheHandler{}, "get,post:Test")
}
