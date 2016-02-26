package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/instagram-beego/database"
)

func main() {
	err := database.InitDBConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	database.RegisterModels()

	// Database alias.
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
