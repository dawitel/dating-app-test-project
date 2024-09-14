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
	// Load configurations for the app
	cfg := config.LoadConfig()
	dsn := cfg.GetDSN()

	// Open a PostgreSQL DB connection
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

	// Initiate the router for the app with error handling
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	router.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Method Not Allowed",
		})
	})

	// Public routes
	router.GET("/api/v1/match/recommendations/:user_id", matchmakingHandler.GetMatchRecommendations)
	router.POST("/api/v1/sign-up", userHandler.CreateUser)

	// Protected routes with JWT authentication
	authRoutes := router.Group("/users")
	authRoutes.Use(authService.AuthMiddleware()) // JWT protection
	{
		authRoutes.POST("/sign-in", userHandler.LoginHandler)
		authRoutes.DELETE("/delete/:user_id", userHandler.DeleteUser)
	}

	router.Run(cfg.Port)
}
