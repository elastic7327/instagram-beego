package routers

import (
	"github.com/astaxie/beego"
	"github.com/instagram-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/api",
		beego.NSRouter(
			"/user",
			&controllers.UserController{},
			"post:CreateUser",
		),
	)

	beego.AddNamespace(ns)
}
