package service

import (
	"errors"
	"net/http"
	"test-matchmaking-app/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	jwtSecret string
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// NewAuthService initializes a new AuthService
func NewAuthService() *AuthService {
	cfg := config.LoadConfig()
	return &AuthService{
		jwtSecret: cfg.JWTSecret,
	}
}

// GenerateToken generates a JWT token for a user
func (s *AuthService) GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates a JWT token and returns the claims if valid
func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// AuthMiddleware is a middleware to protect routes
func (s *AuthService) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Extract token from the "Bearer <token>" scheme
		tokenString := authHeader[len("Bearer "):]

		claims, err := s.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Store user ID in context for protected routes
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
