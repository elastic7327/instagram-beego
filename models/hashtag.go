package models

type Hashtag struct {
	Id     int
	Name   string   `orm:"unique"`
	Photos []*Photo `orm:"reverse(many)"`
}
