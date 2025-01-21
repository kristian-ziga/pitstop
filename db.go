package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Review struct {
	ReviewID    string `json:"reviewId,omitempty"` //intentional to have them omitempty, so I can parse input without errors and after that finds what's missing
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Score       int    `json:"score,omitempty"`
}

func CreateDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to open database:", err)
		return nil, err
	}

	err = db.AutoMigrate(&Review{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
		return nil, err
	}

	return db, nil
}
