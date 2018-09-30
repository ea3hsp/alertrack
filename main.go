package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/ea3hsp/alertrack/handlers"
)

func main() {
	server := http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: handler.NewHandler(),
	}

	log.Printf("[AlertTrack] Starting API HTTP Server. Listening at %q", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Printf("%v", err)
	} else {
		log.Println("Server close !!")
	}
}
