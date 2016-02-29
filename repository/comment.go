package repository

import (
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
)

type CommentRepository struct{}

func (this *CommentRepository) Create(comment *models.Comment) (*models.Comment, error) {
	o := orm.NewOrm()
	_, err := o.Insert(comment)
	return comment, err
}
