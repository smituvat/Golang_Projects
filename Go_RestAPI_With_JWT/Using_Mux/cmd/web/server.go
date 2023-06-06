package main

import (
	"fmt"
	"net/http"

	"example.com/m/v/Go_RestAPI_With_JWT/Using_Mux/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", handlers.SayHello("Vatgal")).Methods("GET")
	r.HandleFunc("/user", handlers.HandleUser).Methods("POST")

	// Server config
	fmt.Println("Server at PORT 8080")
	http.ListenAndServe(":8080", r)
}
