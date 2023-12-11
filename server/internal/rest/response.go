package rest

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(data)
}
