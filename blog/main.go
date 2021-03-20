package main

import (
	"blog/internal/routers"
	"net/http"
)

func main() {
	router := routers.NewRouter()

	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	s.ListenAndServe()
}
