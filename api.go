package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func TestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	}
}

func TestWriteHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Got POST request")
		println(r.Body)
	}
}

func CreateApi(db *gorm.DB, r *mux.Router) {
	r.HandleFunc("/", TestHandler(db)).Methods("GET")
	r.HandleFunc("/write", TestWriteHandler(db)).Methods("POST")
}
