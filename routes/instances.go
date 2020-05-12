package routes

import (
	"github.com/gorilla/mux"
	"github.com/semut-technologies/mongodb/handlers"
)

func LoadInstancesRoutes(r *mux.Router) *mux.Router {

	// Get list of all the available instances
	r.HandleFunc("/instances", handlers.DefaultHandler).Methods("GET")

	// Deploy a new instance
	r.HandleFunc("/instance/deploy", handlers.DefaultHandler).Methods("POST")

	// Get a particular instance
	r.HandleFunc("/instance/get", handlers.DefaultHandler).Methods("GET")

	// Modify a particular instance
	r.HandleFunc("/instance/modify", handlers.DefaultHandler).Methods("PATCH")

	// Delete a particular instance
	r.HandleFunc("/instance/delete", handlers.DefaultHandler).Methods("DELETE")

	return r
}
