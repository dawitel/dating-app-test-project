package api

import (
    "github.com/gin-gonic/gin"
    "test-matchmaking-app/internal/service"
    "net/http"
)

type MatchmakingHandler struct {
    service *service.MatchmakingService
}

func NewMatchmakingHandler(service *service.MatchmakingService) *MatchmakingHandler {
    return &MatchmakingHandler{service: service}
}

func (h *MatchmakingHandler) GetMatchRecommendations(c *gin.Context) {
    userID := c.Param("user_id")

    matches, err := h.service.FindMatchesForUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, matches)
}
