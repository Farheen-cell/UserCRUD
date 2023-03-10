package main

import (
	"fmt"

	"github.com/Farheen-cell/UserCRUD/config"
	"github.com/Farheen-cell/UserCRUD/model"
	"github.com/Farheen-cell/UserCRUD/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql",
		config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status: ", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&model.User{})
	r := routes.SetupRouter()
	//running
	r.Run()
}
