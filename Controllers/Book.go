package Controllers

import (
	"go_restapi/Config"
	"go_restapi/Models"
	"go_restapi/Utility"

	"github.com/gin-gonic/gin"
)

func ListBook(c *gin.Context) {
	var book []Models.Book
	err := Models.GetAllBook(&book)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models.Book
	c.BindJSON(&book)
	err := Models.AddNewBook(&book)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

// GetOneBook godoc
// @summary Get Book
// @description  Get Book by id
// @tags Book
// @security ApiKeyAuth
// @id Book
// @accept json
// @produce json
// @param id path int true "id of Book to be gotten"
// @response 200 {object} Models.Book "OK"
// @Router /book/:id [get]
func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models.Book
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.GetOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models.PutOneBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models.Book
	id := c.Params.ByName("id")
	err := Models.DeleteBook(&book, id)
	if err != nil {
		Utility.JSON(c, 404, book)
	} else {
		Utility.JSON(c, 200, book)
	}
}

func ListBookCustom(c *gin.Context) {
	var result = []struct {
		Id   int
		Name string
	}{}

	rows, err := Config.DB.Raw("select id, name from book").Rows()
	defer rows.Close()
	for rows.Next() {
		Config.DB.ScanRows(rows, &result)
	}
	if err != nil {
		Utility.JSON(c, 500, result)
	} else {
		Utility.JSON(c, 200, result)
	}
}

func GetOneBookCustom(c *gin.Context) {
	id := c.Params.ByName("id")

	var result = struct {
		Id   int
		Name string
	}{}

	rows, err := Config.DB.Raw("select id, name from book where id = ?", id).Rows()
	defer rows.Close()
	for rows.Next() {
		Config.DB.ScanRows(rows, &result)
	}
	if err != nil {
		Utility.JSON(c, 500, result)
	} else {
		Utility.JSON(c, 200, result)
	}
}
