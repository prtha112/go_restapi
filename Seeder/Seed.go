package Seeder

import (
	"go_restapi/Config"
	"go_restapi/Models"
)

func Load() {
	category := Models.Category{
		Name: "Novel",
	}
	Config.DB.Create(&category)
}
