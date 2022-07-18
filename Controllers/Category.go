package Controllers

import (
	"go_restapi/Models"
	"go_restapi/Utility"

	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context) {
	var cate []Models.Category
	err := Models.GetAllCategory(&cate)
	if err != nil {
		Utility.RespondJSON(c, 404, cate)
	} else {
		Utility.RespondJSON(c, 200, cate)
	}
}

func AddNewCategory(c *gin.Context) {
	var cate Models.Category
	c.BindJSON(&cate)
	err := Models.AddNewCategory(&cate)
	if err != nil {
		Utility.RespondJSON(c, 404, cate)
	} else {
		Utility.RespondJSON(c, 200, cate)
	}
}
