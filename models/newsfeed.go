package models

import "github.com/jinzhu/gorm"

type NewsFeed struct {
	gorm.Model
	Title  string `json:"Title"`
	Post   string `json:"Post"`
	Author string `json:"Author"`
}

var NewsFeeds []NewsFeed

func (n *NewsFeed) TableName() string {
	return "Newsfeed"
}

type GetNewsFeedsResponse struct {
	Total int         `json:"Total"`
	Data  *[]NewsFeed `json:"data"`
}

type CreateNewsFeedRequest struct {
	Title  string `json:"title,omitempty" binding:"required"`
	Post   string `json:"post,omitempty" binding:"required"`
	Author string `json:"author,omitempty" binding:"required"`
}

type CreateNewsFeedResponse struct {
	ID uint `json:"id"`
}
