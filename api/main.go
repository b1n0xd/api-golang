package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", getUsers)
	fmt.Println("API online in:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode([]User{
		{
			ID:   1,
			Name: "Leonar",
		},
		{
			ID:   2,
			Name: "Luciano",
		},
	})
}
