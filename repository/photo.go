package repository

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
	"strconv"
)

type PhotoRepository struct{}

func (this *PhotoRepository) GetAll() ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	_, err := o.
		QueryTable(&models.Photo{}).
		RelatedSel().
		All(&photos)

	err = _getPhotosComment(photos)

	return photos, err
}

func (this *PhotoRepository) GetByHashtagName(hashtagName string) ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	_, err := o.
		QueryTable(&models.Photo{}).
		Filter("Hashtags__Hashtag__Name", hashtagName).
		RelatedSel().
		All(&photos)

	err = _getPhotosComment(photos)

	return photos, err
}

func (this *PhotoRepository) GetByUserId(userId int) ([]*models.Photo, error) {
	var photos []*models.Photo
	o := orm.NewOrm()
	ps := o.QueryTable(&models.Photo{})
	_, err := ps.Filter("User", strconv.Itoa(userId)).RelatedSel().All(&photos)

	err = _getPhotosComment(photos)

	return photos, err
}

func _getPhotosComment(photos []*models.Photo) error {
	var err error

	for _, photo := range photos {
		err = _getPhotoComment(photo)
	}

	return err
}

func _getPhotoComment(photo *models.Photo) error {
	o := orm.NewOrm()
	photoIdStr := strconv.Itoa(photo.Id)
	_, err := o.
		QueryTable(&models.Comment{}).
		Filter("Photo__Id", photoIdStr).
		RelatedSel("User").
		OrderBy("-id").
		All(&photo.Comments)

	return err
}
