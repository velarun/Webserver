package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

// AddApproutes will add the routes for the application
func AddApproutes(route *mux.Router) {

	route.HandleFunc("/", renderHome)
	route.HandleFunc("/tray", insertRow).Methods("POST")
	route.HandleFunc("/tray/{id}", getID).Methods("GET")

	fmt.Println("Routes are Loded.")
}
