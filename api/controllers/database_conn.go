package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) ConnectDB(host string, port string, user string, dbname string, password string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	server.DB, err = gorm.Open("postgres", DBURL)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connected!")
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Start(addr string) {
	log.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
