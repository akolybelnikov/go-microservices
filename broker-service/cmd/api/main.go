package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const webPort = "80"

type Config struct {
}

func main() {
	app := Config{}

	log.Printf("Starting server on port %s\n", webPort)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", webPort),
		Handler:           app.routes(),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       2 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
