package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080
	handleRequests()
	log.Printf("Server is starting on port %v", port)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleRequests() {
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
