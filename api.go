package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func TestHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func CreateApi(db *gorm.DB, r *mux.Router) {
	r.HandleFunc("/api/test", TestHandler(db)).Methods("POST")
}
