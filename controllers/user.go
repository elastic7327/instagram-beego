package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/models"
	"github.com/instagram-beego/parser/request"
	"github.com/instagram-beego/parser/response"
	"github.com/instagram-beego/repository"
	"strconv"
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

	_, err := userRepository.Create(&user)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &user
	}

	this.ServeJSON()
}

func (this *UserController) Login() {
	userRepository := repository.UserRepository{}
	loginFormRequest := request.LoginFormRequest{}

	this.ParseForm(&loginFormRequest)

	user, err := userRepository.Login(loginFormRequest.Email, loginFormRequest.Password)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &response.LoginResponse{
			Id:          user.Id,
			DisplayName: user.DisplayName,
			Token:       user.Token,
		}
	}

	this.ServeJSON()
}

func (this *UserController) GetById() {
	userRepository := repository.UserRepository{}
	userId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	}

	user, err := userRepository.GetById(userId)

	// TODO how to disable token in json
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &user
	}

	this.ServeJSON()
}

func (this *UserController) Update() {
	userRepository := repository.UserRepository{}
	updateFormRequest := request.UpdateFormRequest{}

	token := this.Ctx.Input.Header("token")
	this.ParseForm(&updateFormRequest)

	user, err := userRepository.GetByToken(token)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		if updateFormRequest.DisplayName != "" {
			user.DisplayName = updateFormRequest.DisplayName
		}
		if updateFormRequest.Email != "" {
			user.Email = updateFormRequest.Email
		}

		err = userRepository.Update(&user)

		if err != nil {
			this.Ctx.Output.SetStatus(400)
			this.Data["json"] = response.ErrorResponse{
				ExitCode: 1,
				Message:  err.Error(),
			}
		} else {
			this.Data["json"] = &user
		}
	}

	this.ServeJSON()
}
