package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db, err := CreateDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	r := mux.NewRouter()

	CreateApi(db, r)

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "XMLHttpRequest", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
