package Models

import (
	"go_restapi/Config"

	_ "gorm.io/driver/sqlite"
)

type Category struct {
	// gorm.Model
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `gorm:"not null" json:"name"`
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
