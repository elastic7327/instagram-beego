package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/instagram-beego/database"
	_ "github.com/instagram-beego/routers"
	"github.com/instagram-beego/services"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	services.TestS3()

	err := database.InitDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	database.RegisterModels()
	beego.Run()
}
