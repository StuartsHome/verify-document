package service

import (
	"net/http"
)

func HttpRun() {

	router := BuildRouter()
	server := http.Server{
		Addr:    *listenAddress,
		Handler: router,
	}

	server.ListenAndServe()

}
