package main

import (
	"log"
	"net/http"
	"test-matchmaking-app/config"
	"test-matchmaking-app/internal/api"
	"test-matchmaking-app/internal/repository"
	"test-matchmaking-app/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// load configurations for the app
	cfg := config.LoadConfig()
	dsn := cfg.GetDSN()
	
	// open a postgres db connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initiate all needed services
	userRepo := repository.NewUserRepository(db)
	matchmakingService := service.NewMatchmakingService(userRepo)
	matchmakingHandler := api.NewMatchmakingHandler(matchmakingService, userRepo)
	authService := service.NewAuthService()
	userHandler := api.NewUserHandler(userRepo, authService)

	// initiate the router for the app with error handlig
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method Not Allowed",
		})
	})

	// these routes are made public for simlicity
	router.GET("/api/v1/match/recommendations/:user_id", matchmakingHandler.GetMatchRecommendations)
	router.POST("/api/v1/sign-up", userHandler.CreateUser)

	// protected route for demonstrating the auth service functioality 
	router.POST("/api/v1/sign-in", authService.AuthMiddleware(), userHandler.LoginHandler)
	router.DELETE("/api/v1/delete/:user_id", authService.AuthMiddleware(), userHandler.DeleteUser)

	router.Run(cfg.Port)
}
