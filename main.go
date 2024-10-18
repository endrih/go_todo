package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//router.HandleFunc("/api/todo", ).Methods("GET")
	// OauthGoogle
	router.HandleFunc("/auth/google/login", oauthGoogleLogin)
	router.HandleFunc("/auth/google/callback", oauthGoogleCallback)
	port := "10500"
	server := &http.Server{
		Addr:    fmt.Sprintf(":" + port),
		Handler: router,
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
