package repository

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
)

type HashtagRepository struct{}

func (this *HashtagRepository) Create(
	photo *models.Photo,
	hashtag *models.Hashtag,
) (*models.Hashtag, error) {
	o := orm.NewOrm()
	var err error

	err = o.Read(hashtag, "Name")

	if err != nil {
		_, err = o.Insert(hashtag)
	}

	m2m := o.QueryM2M(photo, "Hashtags")
	_, err = m2m.Add(hashtag)

	return hashtag, err
}
