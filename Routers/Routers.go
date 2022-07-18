package Routers

import (
	"fmt"
	auth "go_restapi/Auth"
	"go_restapi/Controllers"
	"log"
	"os"
	"time"

	middleware "go_restapi/Middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	r := gin.Default()

	var env_auth auth.Env
	env_auth.Apikey = os.Getenv("API_TOKEN")
	env_auth.Jwtexpiretime = 5
	env_auth.Jwtsignature = os.Getenv("JWT_SIGNATURE")

	var env_middleware middleware.Env
	env_middleware.Signature = os.Getenv("JWT_SIGNATURE")

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Ping Pong!!!")
	})
	r.GET("/graceful", func(c *gin.Context) {
		fmt.Println("Graceful in process.....")
		time.Sleep(10000 * time.Millisecond)
		fmt.Println("Graceful shutdown complete!!!.")
		c.String(200, "Graceful shutdown test. !!!")
	})

	r.POST("/login", func(c *gin.Context) {
		auth.LoginHandler(c, env_auth)
	})

	book := r.Group("/book", middleware.AuthorizationMiddleware(env_middleware))
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
