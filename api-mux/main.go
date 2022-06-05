package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	mux.Router := mux.NewRouter()
	mux.Router.HandleFunc("/users", getUsers).Methods("GET")
	fmt.Println("API online in port:8080")
	log.Fatal(http.ListenAndServe(":8080", mux.Router))

}
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
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
		{
			ID:   3,
			Name: "Vieira",
		},
		{
			ID:   4,
			Name: "Showtime",
		},
	})
}

