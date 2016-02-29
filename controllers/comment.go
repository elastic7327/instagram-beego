package controllers

import (
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/models"
	"github.com/instagram-beego/parser/request"
	"github.com/instagram-beego/parser/response"
	"github.com/instagram-beego/repository"
	"regexp"
	"strconv"
	"strings"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Create() {
	commentFormRequest := request.CommentFormRequest{}
	commentRepo := repository.CommentRepository{}
	hashtagRepo := repository.HashtagRepository{}
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

	hashtags := _getHashtagsFromContent(comment.Content)
	photo := models.Photo{
		Id: photoId,
	}

	for _, hashtag := range hashtags {
		hashtagModel := models.Hashtag{
			Name: hashtag,
		}
		_, err = hashtagRepo.Create(&photo, &hashtagModel)
		// fmt.Println(err)
	}

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

func _getHashtagsFromContent(content string) []string {
	rx, _ := regexp.Compile("#(?:[[^]]+]|\\S+)")
	hashtags := rx.FindAllString(content, -1)

	for i := range hashtags {
		hashtags[i] = strings.TrimPrefix(hashtags[i], "#")
	}

	return hashtags
}
