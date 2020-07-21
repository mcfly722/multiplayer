package main

import (
	"log"
	"net/http"
)

func main() {
	var bindingAddress = "localhost:8080"

	http.Handle("/", http.FileServer(http.Dir("./static")))

	log.Printf("starting server at %v", bindingAddress)
	log.Fatal(http.ListenAndServe(bindingAddress, nil))
}
