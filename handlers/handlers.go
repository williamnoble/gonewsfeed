package handlers

import (
	"github.com/gin-gonic/gin"
	"gonewsfeed/models"
	"gonewsfeed/services"
	"net/http"
	"strings"
)

func GetAllNews(s *services.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		response, err := s.NewsFeedService.GetAllFeeds()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func PostNewsFeed(services *services.Services) gin.HandlerFunc {
	return func(context *gin.Context) {
		request := models.CreateNewsFeedRequest{}

		// validate json
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":  "json decoding : " + err.Error(),
				"status": http.StatusBadRequest,
			})
			return
		}

		// check empty strings
		if len(strings.TrimSpace(request.Title)) == 0 {
			request.Title = ""
		}

		if len(strings.TrimSpace(request.Post)) == 0 {
			request.Post = ""
		}

		response, err := services.NewsFeedService.CreateFeed(request)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, response)
	}
}

func NoRoute(c *gin.Context) {
	c.JSON(
		http.StatusNotFound,
		gin.H{"status": http.StatusNotFound, "error": "Not Found"})
}

func GetHealth(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{"status": "up"})
}
