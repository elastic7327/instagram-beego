package repository

import (
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
)

type PhotoRepository struct{}

func (this *PhotoRepository) GetByToken(token string) ([]*models.Photo, error) {
	userRepository := UserRepository{}
	var photos []*models.Photo
	o := orm.NewOrm()
	user, err := userRepository.GetUserByToken(token)

	if err == nil {
		qs := o.QueryTable(&models.Photo{})
		qs.Filter("User", user.Id).RelatedSel().All(&photos)
	}

	return photos, err
}
