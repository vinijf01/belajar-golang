package main

import (
	"belajar-golang-depedency-injection/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		// Handler: router,
		Handler: authMiddleware,
	}

}

func main() {

	// server := InitializedServer()

	// err := server.ListenAndServe()
	// helper.PanicIfError(err)

}
