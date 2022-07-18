package Models

import (
	"go_restapi/Config"

	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name"`
}

func (c *Category) TableName() string {
	return "Category"
}

func GetAllCategory(c *[]Category) (err error) {
	if err = Config.DB.Find(c).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCategory(c *Category) (err error) {
	if err = Config.DB.Create(c).Error; err != nil {
		return err
	}
	return nil
}
