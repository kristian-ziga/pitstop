package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	println("Starting service...")
	db, err := CreateDB()
	if err != nil {
		panic("failed to open database:")
	}
	println("DB ok...")

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	println("Router ok...")

	CreateApi(db, r)
	println("API ok...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
