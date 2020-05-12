package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/semut-technologies/mongodb/handlers"
)

// NewRouter returns a new mux router
func NewRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.DefaultHandler)

	return router

}

// NewHandler returns a cors enabled mux handler
func NewHandler() http.Handler {

	router := NewRouter()

	handler := cors.AllowAll().Handler(router)

	return handler

}
