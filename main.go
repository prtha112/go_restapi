package main

import (
	"fmt"
	"go_restapi/Config"
	"go_restapi/Models"
	"go_restapi/Routers"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	Config.DB.AutoMigrate(&Models.Book{}, &Models.Category{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
