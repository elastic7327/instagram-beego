package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/repository"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) GetByToken() {
	token := this.Ctx.Input.Header("token")
	photoRepository := repository.PhotoRepository{}

	photos, err := photoRepository.GetByToken(token)

	if err != nil {
		fmt.Println(err)
	}

	this.Data["json"] = &photos

	this.ServeJSON()
}
