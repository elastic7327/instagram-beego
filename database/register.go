package database

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/models"
	_ "github.com/lib/pq"
)

func InitDBConnection() error {
	driverName := "postgres"

	orm.Debug = true
	orm.RegisterDriver(driverName, orm.DRPostgres)

	return orm.RegisterDataBase(
		"default",
		driverName,
		getConnectionString(),
	)
}

func RegisterModels() {
	orm.RegisterModel(
		&models.User{},
		&models.Hashtag{},
		&models.Photo{},
		&models.Comment{},
	)
}

func getConnectionString() string {
	host := beego.AppConfig.String("db_host")
	username := beego.AppConfig.String("db_user")
	password := beego.AppConfig.String("db_password")
	sslMode := beego.AppConfig.String("db_sslmode")
	port := beego.AppConfig.String("db_port")
	databaseName := beego.AppConfig.String("db_name")

	connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
	return fmt.Sprintf(connectionStringTemplate,
		host, port, sslMode,
		username, password, databaseName)
}
