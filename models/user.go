package models

import (
	"time"
)

type User struct {
	Id          int
	DisplayName string
	Email       string    `orm:"unique"`
	Password    string    `json:"-"`
	Token       string    `json:"-"`
	Photos      []*Photo  `orm:"reverse(many)"`
	Created     time.Time `orm:"auto_now_add;type(datetime) json:"-"`
	Updated     time.Time `orm:"auto_now;type(datetime)" json:"-"`
}
