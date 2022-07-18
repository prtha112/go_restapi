package main

import (
	"fmt"
	"go_restapi/Config"
	"go_restapi/Models"
	"go_restapi/Routers"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/pseidemann/finish"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	Config.DB.AutoMigrate(&Models.Book{}, &Models.Category{})

	r := Routers.SetupRouter()
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
