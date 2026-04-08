package handlers

import (
	"net/http"
	"os"
	"time"

	"notes-api/config"
	"notes-api/models"
	"notes-api/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req models.RegisterRequest

	// Bind JSON body to struct
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Could not hash password")
		return
	}

	// Create user object
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save to DB
	if err := repository.CreateUser(config.DB, user); err != nil {
		errorResponse(c, http.StatusConflict, "Email already exists")
		return
	}

	successResponse(c, "User registered successfully")
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, "Invalid request")
		return
	}

	// Fetch user from DB by email
	user, err := repository.GetUserByEmail(config.DB, req.Email)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		errorResponse(c, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT token
	token, err := generateToken(user.ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "Could not generate token")
		return
	}

	successResponse(c, gin.H{"token": token})
}

func generateToken(userID int) (string, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"issuedAt": time.Now().Unix(),
	})

	return claims.SignedString([]byte(secret))
}
