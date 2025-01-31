package main

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreateApi(db *gorm.DB, r *mux.Router) {
	r.HandleFunc("/review", ReviewAdd(db)).Methods("POST")
	r.HandleFunc("/review", GetReviews(db)).Methods("GET")
}
