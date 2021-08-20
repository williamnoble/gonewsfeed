package router

import (
	"github.com/gin-gonic/gin"
	"gonewsfeed/handlers"
	"gonewsfeed/services"
)

func InitializeRouter(services *services.Services) *gin.Engine {
	r := gin.Default()
	initializeRoutes(r, services)
	return r
}

func initializeRoutes(r *gin.Engine, s *services.Services) {
	r.GET("/health", handlers.GetHealth)
	r.GET("/newsfeeds", handlers.GetAllNews(s))
	r.POST("/newsfeed", handlers.PostNewsFeed(s))
	r.NoRoute(handlers.NoRoute)

}
