package controllers

import (
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/middlewares"
)

func (s *Server) initializeRoutes() {

	// Home Route

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Product routes
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/product", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PATCH")
	s.Router.HandleFunc("/product/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProduct)).Methods("DELETE")

	//Categories routes
	s.Router.HandleFunc("/category", middlewares.SetMiddlewareJSON(s.CreateCategory)).Methods("POST")
	s.Router.HandleFunc("/category", middlewares.SetMiddlewareJSON(s.GetCategories)).Methods("GET")
	s.Router.HandleFunc("/category/{id}", middlewares.SetMiddlewareJSON(s.GetCategory)).Methods("GET")
	s.Router.HandleFunc("/category/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateCategory))).Methods("PATCH")
	s.Router.HandleFunc("/category/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteCategory)).Methods("DELETE")

	// s.Router.HandleFunc("/g", middlewares.SetMiddlewareJSON(s.BasicQ)).Methods("GET")
}
