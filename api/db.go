package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Review struct {
	ReviewID uuid.UUID `json:"reviewId,omitempty"`
	ReviewersInformation
}

type ReviewersInformation struct {
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

func GetReviews(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reviews []Review
		db.Find(&reviews)
		_ = json.NewEncoder(w).Encode(reviews)
	}
}

func ReviewAdd(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		defer func() {
			_ = r.Body.Close()
		}()

		var reviewInfo ReviewersInformation
		err = json.Unmarshal(body, &reviewInfo)
		if err != nil {
			http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
			return
		}

		if reviewInfo.Name == "" || reviewInfo.Score < 1 || reviewInfo.Score > 5 {
			http.Error(w, `{"message": "Invalid input: Name, Description required. Score must be 1-5"}`, http.StatusBadRequest)
			return
		}

		reviewID := uuid.New()

		review := Review{
			ReviewID:             reviewID,
			ReviewersInformation: reviewInfo,
		}

		if err := db.Create(&review).Error; err != nil {
			http.Error(w, "Error adding review to the database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		response := map[string]interface{}{
			"message":  "Review added successfully",
			"reviewId": review.ReviewID.String(),
		}
		_ = json.NewEncoder(w).Encode(response)
	}
}
