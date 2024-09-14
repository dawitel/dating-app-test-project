package api

import (
	"net/http"
	"test-matchmaking-app/internal/domain"
	"test-matchmaking-app/internal/repository"
	"test-matchmaking-app/internal/service"
	"test-matchmaking-app/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserHandler represents a handler for user-related requests.
type UserHandler struct {
	userRepo    *repository.UserRepository
	authService *service.AuthService
}

// NewUserHandler returns a pointer to the UserHandler.
func NewUserHandler(userRepo *repository.UserRepository, authService *service.AuthService) *UserHandler {
	return &UserHandler{
		userRepo:    userRepo,
		authService: authService,
	}
}

// CreateUser handles the create user/sign-up requests.
// CreateUser handles the create user/sign-up requests.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser domain.User

	// Bind JSON request body to the User struct
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Check if a user with the same name already exists
	existingUser, err := h.userRepo.FindUserByName(newUser.Name)
	if err != nil && err != gorm.ErrRecordNotFound { // If there's an error other than "record not found", return internal server error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing user"})
		return
	}
	if existingUser != nil { // User with the same name exists
		c.JSON(http.StatusConflict, gin.H{"error": "User with this name already exists"})
		return
	}

	// Generate a new UUID for the user if not provided
	if newUser.UserID == "" {
		newUser.UserID = uuid.New().String()
	} else {
		// Validate that user_id is a valid UUID if provided
		if _, err := uuid.Parse(newUser.UserID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format. Must be a valid UUID."})
			return
		}
	}

	// Hash the password before saving it to the database
	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}
	newUser.Password = hashedPassword

	// Save the user to the database
	if err := h.userRepo.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	// Generate a JWT token for the user
	token, err := h.authService.GenerateToken(newUser.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token for the user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": newUser.UserID, "token": token})
}


// DeleteUser handles a delete user request.
func (h *UserHandler) DeleteUser(c *gin.Context) {
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

// LoginHandler handles the login requests.
func (h *UserHandler) LoginHandler(c *gin.Context) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON request body to the login request struct
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Retrieve the user by username
	user, err := h.userRepo.GetUserByName(loginRequest.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Check the password
	if err := utils.CheckPassword(loginRequest.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate a JWT token for the user
	token, err := h.authService.GenerateToken(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token for the user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
