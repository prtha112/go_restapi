package Models

import (
	"fmt"
	"go_restapi/Config"
	"time"

	_ "gorm.io/driver/sqlite"
)

type Book struct {
	// gorm.Model
	ID         uint   `gorm:"primary_key" json:"id"`
	Name       string `json:"name"`
	CategoryID int
	Category   Category  // ถ้าอยากให้เเสดงข้อมูลตอน Query ให้ใส่ Preload
	Author     string    `json:"author"`
	CreatedAt  time.Time `gorm:"default now()"`
}

func (b *Book) TableName() string {
	return "book"
}

func GetAllBook(b *[]Book) (err error) {
	// ถ้าไม่ใส่ Preload จะไม่ Join Data กับ Category
	if err = Config.DB.Preload("Category").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewBook(b *Book) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneBook(b *Book, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteBook(b *Book, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
