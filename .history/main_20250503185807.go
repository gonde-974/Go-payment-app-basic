package main

import (
	"go-payment-app-basik/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
