package api

import (
    "net/http"
    "test-matchmaking-app/internal/domain"
    "test-matchmaking-app/internal/repository"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type UserHandler struct {
    userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
    return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var newUser domain.User

    // Bind JSON request body to the User struct
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
        return
    }

    // Generate a new UUID for the user if not provided
    if newUser.UserID == "" {
        newUser.UserID = uuid.New().String()
    } else {
        // Validate user_id is a valid UUID if provided
        if _, err := uuid.Parse(newUser.UserID); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format. Must be a valid UUID."})
            return
        }
    }

    // Save the user to the database
    if err := h.userRepo.CreateUser(&newUser); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": newUser.UserID})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    // Get the user_id from the URL parameter
    userID := c.Param("user_id")

    // Validate if user_id is a valid UUID
    if _, err := uuid.Parse(userID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format. Must be a valid UUID."})
        return
    }

    // Delete the user using the repository
    if err := h.userRepo.DeleteUser(userID); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete user"})
        return
    }

    // Return success response
    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
