package main

import (
	"log"
	"net/http"
	"os"
	"weather-app/handlers"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	addr := ":8080"
	http.HandleFunc("/weather", handlers.WeatherHandler)

	log.Printf("Server is running on %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}