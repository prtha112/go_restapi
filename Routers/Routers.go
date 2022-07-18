package Routers

import (
	"go_restapi/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	book := r.Group("/book")
	{
		book.GET("", Controllers.ListBook)
		book.POST("", Controllers.AddNewBook)
		book.GET(":id", Controllers.GetOneBook)
		book.PUT(":id", Controllers.PutOneBook)
		book.DELETE(":id", Controllers.DeleteBook)
	}

	category := r.Group("/category")
	{
		category.GET("", Controllers.ListCategory)
		category.POST("", Controllers.AddNewCategory)
	}

	return r
}
