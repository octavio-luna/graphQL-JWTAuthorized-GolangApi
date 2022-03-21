package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/controllers"
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/models"
)

var server = controllers.Server{}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.ConnectDB(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))

	server.Start(":8080")

}

func ConfigDB() {
	server.DB.CreateTable(&models.User{}, &models.Product{}, &models.Categories{})
}
