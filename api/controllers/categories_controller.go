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

func (server *Server) CreateCategory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
	}
	category := models.Categories{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	category.Prepare()
	err = category.Validate("")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	categoryCreated, err := category.SaveCategory(server.DB)

	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Id", fmt.Sprintf("%d", categoryCreated.ID))
	responses.Write(w, http.StatusCreated, categoryCreated)
}

func (server *Server) GetCategories(w http.ResponseWriter, r *http.Request) {

	category := models.Categories{}

	categories, err := category.FindAllCategories(server.DB)
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	responses.Write(w, http.StatusOK, categories)
}

func (server *Server) GetCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	category := models.Categories{}
	categoryGot, err := category.FindCategoryByID(server.DB, uint32(uid))
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	responses.Write(w, http.StatusOK, categoryGot)
}

func (server *Server) UpdateCategory(w http.ResponseWriter, r *http.Request) {

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
	category := models.Categories{}
	err = json.Unmarshal(body, &category)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	category.Prepare()
	err = category.Validate("update")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}
	updatedCategory, err := category.UpdateCategory(server.DB, uint32(uid))
	if err != nil {
		responses.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	responses.Write(w, http.StatusOK, updatedCategory)
}

func (server *Server) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	product := models.Product{}

	err = errors.New("Category has associated products, delete products first")
	products, _ := product.FindAllProducts(server.DB)
	for i := range *products {
		fmt.Println(uint32(uid), (*products)[i].ID)
		if uint32(uid) == (*products)[i].ID {
			responses.WriteError(w, http.StatusBadRequest, err)
			return
		}
	}
	category := models.Categories{}
	_, err = category.DeleteACategory(server.DB, uint32(uid))
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.Write(w, http.StatusOK, "Deleted")
}
