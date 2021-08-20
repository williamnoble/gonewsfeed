package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gonewsfeed/config"
	"gonewsfeed/router"
	"gonewsfeed/services"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type App struct {
	Config   config.Config
	Router   *gin.Engine
	Services *services.Services
}

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("error when loading configuration")
	}

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal("Failed to open database")
	}
	s := services.NewServices(db)
	r := router.InitializeRouter(s)

	app := &App{
		Config:   cfg,
		Router:   r,
		Services: s,
	}

	err = app.Router.Run()
	if err != nil {
		log.Fatal("Failed to start Gin")
	}
}

func openDB(_ config.Config) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("failed to open db")
		return nil, errors.New("encountered an error generating postgres")
	}

	return db, nil
}
