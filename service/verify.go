package service

import (
	"encoding/json"
	"net/http"
)

func verify(w http.ResponseWriter, r *http.Request) {
	words := "hello world"

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&words)
}
