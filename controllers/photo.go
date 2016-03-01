package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/models"
	"github.com/instagram-beego/parser/response"
	"github.com/instagram-beego/repository"
	"github.com/instagram-beego/services"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"time"
)

func init() {
	godotenv.Load()
}

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) Create() {
	userRepo := repository.UserRepository{}
	photoRepo := repository.PhotoRepository{}

	token := this.Ctx.Input.Header("token")
	user, getUserErr := userRepo.GetByToken(token)

	if getUserErr != nil {
		fmt.Println("Get user by token failed: ", getUserErr.Error())
		return
	}

	file, header, _ := this.GetFile("file")
	bits := strings.Split(header.Filename, ".")
	fileExt := bits[len(bits)-1]
	fileName := strconv.Itoa(int(time.Now().Unix())) + "." + fileExt

	uploadS3Err := services.UploadS3(file, fileName)

	if uploadS3Err != nil {
		fmt.Println("Upload failed: ", uploadS3Err.Error())
		return
	}

	photo := models.Photo{
		Url:  "http://" + os.Getenv("BUCKET") + ".s3-website-" + os.Getenv("AWS_REGION") + ".amazonaws.com/" + fileName,
		User: &user,
	}

	_, createPhotoErr := photoRepo.Create(&photo)

	if createPhotoErr != nil {
		fmt.Println("Create photo failed: ", createPhotoErr.Error())
		return
	}

	this.Data["json"] = &photo

	this.ServeJSON()
}

func (this *PhotoController) GetAll() {
	var photos []*models.Photo
	var err error
	photoRepository := repository.PhotoRepository{}
	q := this.Ctx.Input.Query("tag")

	if q == "" {
		photos, err = photoRepository.GetAll()
	} else {
		photos, err = photoRepository.GetByHashtagName(q)
	}

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = photos
	}

	this.ServeJSON()
}

func (this *PhotoController) GetByUserId() {
	photoRepository := repository.PhotoRepository{}
	userId, err := strconv.Atoi(this.Ctx.Input.Param(":id"))

	photos, err := photoRepository.GetByUserId(userId)

	if err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Data["json"] = &response.ErrorResponse{
			ExitCode: 1,
			Message:  err.Error(),
		}
	} else {
		this.Data["json"] = &photos
	}

	this.ServeJSON()
}
