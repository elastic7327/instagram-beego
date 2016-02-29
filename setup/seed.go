package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/database"
	"github.com/instagram-beego/models"
)

func main() {
	err := database.InitDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	database.RegisterModels()

	o := orm.NewOrm()

	users := []*models.User{
		&models.User{
			DisplayName: "Tri Truong",
			Email:       "tri.itvn@gmail.com",
			Password:    "123456",
		},
		&models.User{
			DisplayName: "Lan Xuan",
			Email:       "lan@gmail.com",
			Password:    "12345",
			Photos: []*models.Photo{
				&models.Photo{
					Url: "url-3",
				},
			},
		},
	}

	photos := []*models.Photo{
		&models.Photo{
			Url:  "url-1",
			User: users[0],
		},
		&models.Photo{
			Url:  "url-2",
			User: users[0],
		},
		&models.Photo{
			Url:  "url-3",
			User: users[1],
		},
	}

	comments := []*models.Comment{
		&models.Comment{
			Content: "Comment 1",
			Photo:   photos[0],
			User:    users[1],
		},
		&models.Comment{
			Content: "Comment 2",
			Photo:   photos[0],
			User:    users[0],
		},
		&models.Comment{
			Content: "Comment 3",
			Photo:   photos[0],
			User:    users[1],
		},
	}

	for _, user := range users {
		o.Insert(user)
	}

	for _, photo := range photos {
		o.Insert(photo)
	}

	for _, comment := range comments {
		o.Insert(comment)
	}
}
