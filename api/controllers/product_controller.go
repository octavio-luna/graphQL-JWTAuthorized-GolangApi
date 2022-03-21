package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/models"
	responses "github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/utils/response_handlers"
)

func (server *Server) CreateProduct(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
	}
	product := models.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product.Prepare(product.Category)
	err = product.Validate()
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	category := models.Categories{}
	categories, _ := category.FindAllCategories(server.DB)
	err = errors.New("Category ID not found")
	for i := range *categories {
		if product.Category == (*categories)[i].ID {
			err = nil
		}
	}
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	productCreated, err := product.SaveProduct(server.DB)

	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Id", fmt.Sprintf("%d", productCreated.ID))
	responses.Write(w, http.StatusCreated, productCreated)
}

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {

	product := models.Product{}

	products, err := product.FindAllProducts(server.DB)
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	responses.Write(w, http.StatusOK, products)
}

func (server *Server) GetProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product := models.Product{}
	productGotten, err := product.FindProductByID(server.DB, uint32(uid))
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	responses.Write(w, http.StatusOK, productGotten)
}

func (server *Server) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	product := models.Product{}
	err = json.Unmarshal(body, &product)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product.Prepare(product.Category)
	err = product.Validate()
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	category := models.Categories{}
	categories, _ := category.FindAllCategories(server.DB)
	err = errors.New("Category ID not found")
	for i := range *categories {
		if product.Category == (*categories)[i].ID {
			err = nil
		}
	}
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedProduct, err := product.UpdateAProduct(server.DB, uint32(uid))
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	responses.Write(w, http.StatusOK, updatedProduct)
}

func (server *Server) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	product := models.Product{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	_, err = product.DeleteAProduct(server.DB, uint32(uid))
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.Write(w, http.StatusOK, "Deleted")
}
