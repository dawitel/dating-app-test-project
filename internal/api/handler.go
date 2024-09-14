package api

import (
	"net/http"
	"strconv"
	"test-matchmaking-app/internal/repository"
	"test-matchmaking-app/internal/service"

	"github.com/gin-gonic/gin"
)

// MatchmakingHandler represets the handler for the metchmaking process.
type MatchmakingHandler struct {
	service  *service.MatchmakingService
	UserRepo *repository.UserRepository
}

// NewMatchmakingHandler returs a pointer to a new matchmakig hadler.
func NewMatchmakingHandler(service *service.MatchmakingService, userRepo *repository.UserRepository) *MatchmakingHandler {
	return &MatchmakingHandler{service: service, UserRepo: userRepo}
}

// GetMatchRecommendations retrieves potential matches for the user with pagination
func (h *MatchmakingHandler) GetMatchRecommendations(c *gin.Context) {
	// Get userID from context (after middleware authentication)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Retrieve the user from the database
	user, err := h.UserRepo.GetUserByID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
		return
	}

	// Pagination parameters
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // Default limit
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0 // Default offset
	}

	// Fetch matching users based on preferences, interests, and activity
	matches, err := h.UserRepo.GetMatchesForUser(user, limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving matches"})
		return
	}

	// Return the matches as a response with pagination
	c.JSON(http.StatusOK, gin.H{
		"matches": matches,
		"limit":   limit,
		"offset":  offset,
	})
}
