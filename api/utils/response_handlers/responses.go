package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Write(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		Write(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	Write(w, http.StatusBadRequest, nil)
}
