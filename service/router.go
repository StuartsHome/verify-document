package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

// var verifyDocumentService = verify.NewVerifyDocumentService()

func BuildRouter() http.Handler {

	// service := VerifyService{verifies: []verify.Verify{
	// 	verify.NewVerifyReportService()}}

	processService := ProcessService{processors: []Process{
		NewProcessData(),
	}}

	r := mux.NewRouter()

	r.HandleFunc("/verify", processService.VerifyHandler).Methods("GET")
	return r
}
