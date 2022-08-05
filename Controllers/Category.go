package Controllers

import (
	"go_restapi/Models"
	"go_restapi/Utility"

	"github.com/gin-gonic/gin"
)

// ListCategory godoc
// @summary Get Category
// @description  Get Category by id
// @tags Category
// @security ApiKeyAuth
// @id ListCategory
// @accept json
// @produce json
// @param id path int true "id of Category to be gotten"
// @response 200 {object} Models.Category "OK"
// @Router /api/v1/category/:id [get]
func ListCategory(c *gin.Context) {
	var cate []Models.Category
	err := Models.GetAllCategory(&cate)
	if err != nil {
		Utility.JSON(c, 404, cate)
	} else {
		Utility.JSON(c, 200, cate)
	}
}

func AddNewCategory(c *gin.Context) {
	var cate Models.Category
	c.BindJSON(&cate)
	err := Models.AddNewCategory(&cate)
	if err != nil {
		Utility.JSON(c, 404, cate)
	} else {
		Utility.JSON(c, 200, cate)
	}
}

func GetOneCategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var cate Models.Category
	err := Models.GetOneCategory(&cate, id)
	if err != nil {
		Utility.JSON(c, 404, cate)
	} else {
		Utility.JSON(c, 200, cate)
	}
}

func PutOneCategory(c *gin.Context) {
	var cate Models.Category
	id := c.Params.ByName("id")
	err := Models.GetOneCategory(&cate, id)
	if err != nil {
		Utility.JSON(c, 404, cate)
	}
	c.BindJSON(&cate)
	err = Models.PutOneCategory(&cate, id)
	if err != nil {
		Utility.JSON(c, 404, cate)
	} else {
		Utility.JSON(c, 200, cate)
	}
}

func DeleteCategory(c *gin.Context) {
	var cate Models.Category
	id := c.Params.ByName("id")
	err := Models.DeleteCategory(&cate, id)
	if err != nil {
		Utility.JSON(c, 404, cate)
	} else {
		Utility.JSON(c, 200, cate)
	}
}
