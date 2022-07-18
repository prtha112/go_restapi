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

func GetOneCategory(c *Category, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCategory(c *Category, id string) (err error) {
	Config.DB.Save(c)
	return nil
}

func DeleteCategory(c *Category, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(c)
	return nil
}
