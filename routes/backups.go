package routes

import (
	"github.com/gorilla/mux"
	"github.com/semut-technologies/mongodb/handlers"
)

func LoadBackupRoutes(r *mux.Router) *mux.Router {

	// Get list of all the available backups
	r.HandleFunc("/backup", handlers.DefaultHandler).Methods("GET")

	// Trigger a new backup job
	r.HandleFunc("/backup", handlers.DefaultHandler).Methods("POST")

	// Get a particular backup snapshot with snapshotid
	r.HandleFunc("/backup/{id}", handlers.DefaultHandler).Methods("GET")

	// Change expiration time of a particular backup snapshot with snapshotid
	r.HandleFunc("/backup/{id}", handlers.DefaultHandler).Methods("PATCH")

	// Delete a particular backup snapshot with snapshotid
	r.HandleFunc("/backup/{id}", handlers.DefaultHandler).Methods("DELETE")

	// Get the schedule policy of backup
	r.HandleFunc("/backup/schedule", handlers.DefaultHandler).Methods("GET")

	// Change the schedule policy of backup
	r.HandleFunc("/backup/schedule", handlers.DefaultHandler).Methods("PATCH")

	// Get list of all the restore jobs
	r.HandleFunc("/restore", handlers.DefaultHandler).Methods("GET")

	// Trigger a new restore process with snapshot and instance id
	r.HandleFunc("/restore", handlers.DefaultHandler).Methods("POST")

	// Get the details of a particular restore job with restore id
	r.HandleFunc("/restore/{id}", handlers.DefaultHandler).Methods("GET")

	return r
}
