package services

import (
	"gonewsfeed/models"
	"gorm.io/gorm"
	"log"
)

type NewsFeedService struct {
	DB *gorm.DB
}

func NewNewsFeed(db *gorm.DB) (*NewsFeedService, error) {
	err := db.AutoMigrate(&models.NewsFeed{})
	if err != nil {
		log.Fatal("Encountered fatal error when attempting to automigrate")
	}
	return &NewsFeedService{DB: db}, nil
}

func (n NewsFeedService) GetAllFeeds() (*models.GetNewsFeedsResponse, error) {
	d := n.DB.Find(&models.NewsFeeds)
	if d.Error != nil {
		return nil, d.Error
	}

	response := models.GetNewsFeedsResponse{}
	response.Total = len(models.NewsFeeds)
	response.Data = &models.NewsFeeds

	return &response, nil
}

func (n NewsFeedService) CreateFeed(in models.CreateNewsFeedRequest) (*models.CreateNewsFeedResponse, error) {
	m := models.NewsFeed{
		Title:  in.Title,
		Post:   in.Post,
		Author: in.Author,
	}

	feed := models.NewsFeed{}
	d := n.DB.Create(&m).Scan(&feed)

	if d.Error != nil {
		return nil, d.Error
	}

	return &models.CreateNewsFeedResponse{ID: feed.ID}, nil
}
