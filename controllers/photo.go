package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/repository"
)

type PhotoController struct {
	beego.Controller
}

func (this *PhotoController) GetAll() {
	photoRepository := repository.PhotoRepository{}
	photos, err := photoRepository.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	this.Data["json"] = &photos

	this.ServeJSON()
}
