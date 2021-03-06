package repository

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
	"time"
)

type UserRepository struct{}

func (this *UserRepository) Create(user *models.User) (*models.User, error) {
	o := orm.NewOrm()
	_, err := o.Insert(user)
	return user, err
}

func (this *UserRepository) Login(email string, password string) (models.User, error) {
	o := orm.NewOrm()
	user := models.User{
		Email:    email,
		Password: password,
	}
	err := o.Read(&user, "Email", "Password")

	fmt.Println(user)

	// Update Token
	if err == nil {
		user.Token = generateToken()
		err = this.Update(&user)
	}

	return user, err
}

func (this *UserRepository) Update(user *models.User) error {
	o := orm.NewOrm()
	err := o.Read(&models.User{Id: user.Id})

	if err == nil {
		_, err = o.Update(user)
	}

	return err
}

func (this *UserRepository) GetByToken(token string) (models.User, error) {
	o := orm.NewOrm()
	user := models.User{
		Token: token,
	}
	err := o.Read(&user, "Token")
	return user, err
}

func (this *UserRepository) GetById(id int) (models.User, error) {
	o := orm.NewOrm()
	user := models.User{Id: id}
	err := o.Read(&user)
	return user, err
}

func (this *UserRepository) ValidateToken(token string) (models.User, error) {
	o := orm.NewOrm()
	user := models.User{
		Token: token,
	}

	err := o.Read(&user)
	return user, err
}

func generateToken() string {
	timestamp := fmt.Sprint(time.Now().Unix())
	// TODO add random string
	sum := md5.Sum([]byte(timestamp))
	return fmt.Sprintf("%x", sum)
}
