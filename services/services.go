package services

import (
	"gorm.io/gorm"
	"log"
)

type Services struct {
	NewsFeedService *NewsFeedService
}

func NewServices(db *gorm.DB) *Services {
	// Need to instantiate here rather than simply passing NewsFeedService{DB:db} to allow for automigrate
	newsFeedService, err := NewNewsFeed(db)
	if err != nil {
		log.Fatal("Failed to load NewsFeed Service")
	}

	return &Services{
		NewsFeedService: newsFeedService,
	}
}
