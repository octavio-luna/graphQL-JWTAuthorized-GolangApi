package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("This is the home page. Please log in")
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}
