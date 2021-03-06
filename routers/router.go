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
			"/user/:id/photos",
			&controllers.PhotoController{},
			"get:GetByUserId",
		),
		beego.NSRouter(
			"/user/:id",
			&controllers.UserController{},
			"get:GetById;post:Update",
		),
		beego.NSRouter(
			"/photo/:photoId/comment",
			&controllers.CommentController{},
			"post:Create",
		),
		beego.NSRouter(
			"/photo",
			&controllers.PhotoController{},
			"post:Create",
		),
		beego.NSRouter(
			"/photos",
			&controllers.PhotoController{},
			"get:GetAll",
		),
	)

	beego.AddNamespace(ns)
}
