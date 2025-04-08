package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker on port %s.", port)

	// define server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	// start server
	if err := srv.ListenAndServe(); err != nil {
		log.Panic()
	}
}
