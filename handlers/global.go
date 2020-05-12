package handlers

import (
	"fmt"
	"net/http"
)

// DefaultHandler returns hello world
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello world!")
}
