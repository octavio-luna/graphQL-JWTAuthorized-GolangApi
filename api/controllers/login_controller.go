package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/auth"
	"github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/models"
	responses "github.com/octavio-luna/graphQL-JWTAuthorized-GolangApi/api/utils/response_handlers"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.WriteError(w, http.StatusBadRequest, err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		responses.WriteError(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.Write(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
