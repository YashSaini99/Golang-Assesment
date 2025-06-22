package controllers

import (
	"net/http"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service" // TODO: Replace with the correct import path, e.g., "github.com/yourusername/yourrepo/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Login godoc
// @Summary Login
// @Description Login for both receptionist and doctor, returns JWT token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param credentials body object{username=string,password=string} true "User credentials"
// @Success 200 {object} map[string]string "token"
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		userRepo := &repository.UserRepository{DB: db}
		token, err := service.AuthenticateUser(creds.Username, creds.Password, userRepo)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// Register godoc
// @Summary Register user
// @Description Register a new user (receptionist or doctor)
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		// Hash the password before saving
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
		repo := &repository.UserRepository{DB: db}
		if err := repo.Create(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
		c.JSON(http.StatusCreated, user)
	}
}
