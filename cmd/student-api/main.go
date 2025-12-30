package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dangbros/student-api/internal/config"
)

func main() {
	fmt.Println("Welcome to Student API!")
	// load config
	cfg := config.MustLoad()
	// database setup
	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})
	// setup server

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	fmt.Printf("server started %s", cfg.HTTPServer.Address)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("failed to start the server")
	}

}
