package models

type Hashtag struct {
	Id     int
	Name   string
	Photos []*Photo `orm:"reverse(many)"`
}
