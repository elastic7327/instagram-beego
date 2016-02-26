package models

import (
	"time"
)

type Comment struct {
	Id      int
	User    *User  `orm:"rel(fk)"`
	Photo   *Photo `orm:"rel(fk)"`
	Content string
	Created time.Time `orm:"auto_now_add;type(datetime)"`
}
