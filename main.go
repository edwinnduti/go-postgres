package main

import (
	"github.com/edwinnduti/postgres-login/router"
	"log"
	"os"
	"net/http"
)

func main() {
	// call router
	r := router.Router()

	// establish portnumber
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8030"
	}

	// set server
	server := &http.Server{
		Handler: r, // n
		Addr:    ":" + Port,
	}

	// log server output
	log.Printf("Listening on PORT: %s", Port)
	server.ListenAndServe()
}
