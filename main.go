package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/semut-technologies/mongodb/routes"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	PORT := ":" + os.Getenv("PORT")

	handler := routes.NewHandler()

	log.Println("starting server on", PORT)
	log.Fatal(http.ListenAndServe(PORT, handler))
}
