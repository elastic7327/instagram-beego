package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/database"
	_ "github.com/instagram-beego/routers"
)

func main() {
	err := database.InitDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	database.RegisterModels()
	beego.Run()
}
