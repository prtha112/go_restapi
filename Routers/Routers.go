package Routers

import (
	"go_restapi/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Ping Pong!!!")
	})

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

	subcategory := category.Group("/sub")
	{
		subcategory.GET("/", func(c *gin.Context) {
			c.String(200, "Ping Pong!!!")
		})
	}

	return r
}
