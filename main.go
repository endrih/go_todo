package main

import (
	"database/sql"
	"endrih/go_todo/auth"
	"endrih/go_todo/config"
	"endrih/go_todo/data"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	DB       sql.DB
	Session  *sessions.Store
	Config   *config.AppConfig
}

var App *Application

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db := data.Initialize()
	config := config.Initialize()

	App = &Application{
		infoLog:  infoLog,
		errorLog: errorLog,
		DB:       *db,
		Config:   config,
	}
	auth.NewAuth()
	router := mux.NewRouter()
	//router.HandleFunc("/api/todo", ).Methods("GET")
	// OauthGoogle
	router.HandleFunc("/auth/{provider}/login", auth.OauthGoogleLogin)
	router.HandleFunc("/auth/{provider}/logout", auth.OauthGoogleLogout)
	router.HandleFunc("/auth/{provider}/callback", auth.OauthGoogleCallback)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/todoapp/dist/todoapp/browser")))

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
