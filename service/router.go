package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func BuildRouter() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/verify", verify).Methods("GET")
	return r
}
