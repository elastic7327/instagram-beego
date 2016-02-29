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

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Create() {
	commentFormRequest := request.CommentFormRequest{}
	commentRepo := repository.CommentRepository{}
	userRepo := repository.UserRepository{}
	var err error
	var user models.User

	token := this.Ctx.Input.Header("token")
	photoId, _ := strconv.Atoi(this.Ctx.Input.Param(":photoId"))
	this.ParseForm(&commentFormRequest)
	user, err = userRepo.GetByToken(token)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  "User not found (wrong token)",
		}

		this.ServeJSON()
		return
	}

	comment := models.Comment{}
	comment.Content = commentFormRequest.Content
	comment.Photo = &models.Photo{
		Id: photoId,
	}
	comment.User = &user

	this.Data["json"] = &comment
	_, err = commentRepo.Create(&comment)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &comment
	}

	this.ServeJSON()
}
