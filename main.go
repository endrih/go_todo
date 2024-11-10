package main

import (
	"endrih/go_todo/application"
	"endrih/go_todo/auth"
	"fmt"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	application.App.Initialize()
	auth.NewAuth()
	router := mux.NewRouter()
	//router.HandleFunc("/api/todo", ).Methods("GET")
	// OauthGoogle
	router.HandleFunc("/auth/{provider}/login", auth.OauthGoogleLogin)
	router.HandleFunc("/auth/{provider}/logout", auth.OauthGoogleLogout)
	router.HandleFunc("/auth/{provider}/callback", auth.OauthGoogleCallback)

	//Angular frontend
	// Serve static assets directly.
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir, file := path.Split(strings.Split(r.RequestURI, "?")[0])
		ext := filepath.Ext(file)
		application.App.InfoLog.Println("request " + r.RequestURI)
		application.App.InfoLog.Println("dir " + dir)
		application.App.InfoLog.Println("file " + file)
		application.App.InfoLog.Println("ext " + ext)
		if file == "" || ext == "" {
			application.App.InfoLog.Println("serving index.html")
			http.ServeFile(w, r, "./frontend/todoapp/dist/todoapp/browser/index.html")
			return
		} else {
			application.App.InfoLog.Println("serving " + path.Join(dir, file))
			http.ServeFile(w, r, "./frontend/todoapp/dist/todoapp/browser/"+path.Join(dir, file))
			return
		}
	})

	// Catch-all: Serve our JavaScript application's entry-point (index.html).
	// router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./frontend/todoapp/dist/todoapp/browser/index.html")
	// })
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
