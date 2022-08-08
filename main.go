package main

import (
	"fmt"
	auth "go_restapi/Auth"
	"go_restapi/Config"
	middleware "go_restapi/Middleware"
	"go_restapi/Models"
	"go_restapi/Routers"
	"go_restapi/Seeder"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/pseidemann/finish"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "go_restapi/Document"
)

var err error

func main() {
	var err error
	var env_auth auth.Env
	var env_middleware middleware.Env

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	Config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	env_auth.Apikey = os.Getenv("API_TOKEN")
	env_auth.Jwtexpiretime, err = strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
	if err != nil {
		env_auth.Jwtexpiretime = 5 // minute
	}
	env_auth.Jwtsignature = os.Getenv("JWT_SIGNATURE")
	env_middleware.Signature = os.Getenv("JWT_SIGNATURE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	Config.DB.AutoMigrate(&Models.Book{}, &Models.Category{}) // Create table
	Seeder.Load()                                             // Mockup data to table

	r := Routers.SetupRouter(env_auth, env_middleware)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// running
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()
}
