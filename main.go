package main

import (
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api"
	// "github.com/graphql-go/graphql"
)

func main() {
	//If the first time running the api, uncomment configDB to migrate the models
	// api.ConfigDB()
	api.Run()
}
