package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Set response as JSON
	w.Header().Set("Content-Type", "application/json")

	// Decode body to put into object
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)

	json.NewEncoder(w).Encode(user)
}

func SayHello(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Say Hello to %s", name)
	}
}

/*
func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Say Hello to All")
 }
*/
