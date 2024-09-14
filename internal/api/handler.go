package api

import (
	"fmt"
	"net/http"
	"strconv"
	"test-matchmaking-app/internal/repository"
	"test-matchmaking-app/internal/service"

	"github.com/gin-gonic/gin"
)

// MatchmakingHandler represents the handler for the matchmaking process.
type MatchmakingHandler struct {
	service  *service.MatchmakingService
	UserRepo *repository.UserRepository
}

// NewMatchmakingHandler returns a pointer to a new matchmaking handler.
func NewMatchmakingHandler(service *service.MatchmakingService, userRepo *repository.UserRepository) *MatchmakingHandler {
	return &MatchmakingHandler{service: service, UserRepo: userRepo}
}

func (h *MatchmakingHandler) GetMatchRecommendations(c *gin.Context) {
	userID := c.Param("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := h.UserRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // Default to the first page
	}

	pageSize := 10 // Define the number of items per page
	offset := (page - 1) * pageSize

	fmt.Printf("Page: %d, PageSize: %d, Offset: %d\n", page, pageSize, offset) // Debug log

	matches, err := h.UserRepo.GetMatchesForUser(user, pageSize, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving matches"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"matches": matches,
		"page":    page,
		"pageSize": pageSize,
	})
}
