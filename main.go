package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", getBLock).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/", newBlock).Methods("POST")

	log.Println("Server is running on port: 8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
