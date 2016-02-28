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
		beego.NSRouter(
			"/user/login",
			&controllers.UserController{},
			"get:Login",
		),
		beego.NSRouter(
			"/user/:id",
			&controllers.UserController{},
			"get:GetById;post:Update",
		),
		beego.NSRouter(
			"/photos",
			&controllers.PhotoController{},
			"get:GetByToken",
		),
	)

	beego.AddNamespace(ns)
}
