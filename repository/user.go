package repository

import (
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
)

type UserRepository struct{}

func (this *UserRepository) Create(user *models.User) (*models.User, error) {
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return user, err
}
