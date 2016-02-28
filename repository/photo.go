package repository

import (
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
)

type PhotoRepository struct{}

func (this *PhotoRepository) GetAll() ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	_, err := o.QueryTable(&models.Photo{}).RelatedSel().All(&photos)

	return photos, err
}
