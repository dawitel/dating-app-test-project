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
    dsn, err := config.GetDSN()
    if err != nil {
        log.Fatal("Missing Environment variables:", err)
    }
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    userRepo := repository.NewUserRepository(db)
    matchmakingService := service.NewMatchmakingService(userRepo)
    matchmakingHandler := api.NewMatchmakingHandler(matchmakingService)
    userHandler := api.NewUserHandler(userRepo)
    
    router := gin.Default()
    router.HandleMethodNotAllowed = true
    router.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{
            "error": "Method Not Allowed",
        })
    })
    router.GET("/api/match/recommendations/:user_id", matchmakingHandler.GetMatchRecommendations)
    router.POST("/api/create-user", userHandler.CreateUser)
    router.DELETE("/api/delete/:user_id", userHandler.DeleteUser)

    router.Run(":8080")
}
