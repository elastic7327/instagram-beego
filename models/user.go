package models

import (
	"time"
)

type User struct {
	Id          int
	DisplayName string
	Email       string `orm:"unique"`
	Password    string `json:""`
	Token       string
	Photos      []*Photo  `orm:"reverse(many)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
}
