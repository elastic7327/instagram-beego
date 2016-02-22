package routers

import (
	"github.com/instagram-beego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
