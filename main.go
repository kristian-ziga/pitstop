package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	println("Starting service...")
	db, err := CreateDB()
	if err != nil {
		panic("failed to open database:")
	}
	println("DB ok...")

	r := mux.NewRouter()

	println("Router ok...")

	CreateApi(db, r)
	println("API ok...")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "XMLHttpRequest", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8000"+os.Getenv("PORT"), handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
