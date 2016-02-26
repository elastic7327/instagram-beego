package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/models"
	"github.com/instagram-beego/parser/request"
	"github.com/instagram-beego/repository"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) CreateUser() {
	userRepository := repository.UserRepository{}
	registerFormRequest := request.RegisterFormRequest{}

	this.ParseForm(&registerFormRequest)
	user := models.User{
		DisplayName: registerFormRequest.DisplayName,
		Email:       registerFormRequest.Email,
		Password:    registerFormRequest.Passsword,
	}
	userRepository.Create(&user)

	this.Ctx.Output.SetStatus(400)

	this.Data["json"] = &user
	this.ServeJSON()
}
