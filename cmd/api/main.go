package main

import (
	"log"

	httpserver "distributed-url-shortener/internal/http"
)

func main() {
	log.Println("Distributed URL Shortener starting...")

	router := httpserver.NewRouter()
	server := httpserver.NewServer(":8080", router)

	server.Start()
}
