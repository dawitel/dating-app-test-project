package api

import (
	"net/http"
	"strconv"
	"test-matchmaking-app/internal/domain"
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

    // Retrieve user by ID
    user, err := h.UserRepo.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving user"})
        return
    }

    // Check if the preferences fields are initialized
    if user.Preferences.MinAge == 0 || user.Preferences.MaxAge == 0 || user.Preferences.Gender == "" || user.Preferences.MaxDistance == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User preferences not set or incomplete"})
        return
    }

    // Parse pagination info from query
    pageStr := c.Query("page")
    page, err := strconv.Atoi(pageStr)
    if err != nil || page <= 0 {
        page = 1 // Default to page 1
    }

    pageSize := 10 // Items per page
    offset := (page - 1) * pageSize

    // Fetch matches for user based on preferences
    matches, totalMatches, err := h.service.GetMatchesForUser(user, pageSize, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving matches"})
        return
    }

    // If no matches or page exceeds available matches, return empty response
    if len(matches) == 0 || offset >= totalMatches {
        c.JSON(http.StatusOK, gin.H{
            "matches":     []domain.User{},
            "page":        page,
            "pageSize":    pageSize,
            "totalMatches": totalMatches,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "matches":     matches,
        "page":        page,
        "pageSize":    pageSize,
        "totalMatches": totalMatches,
    })
}
