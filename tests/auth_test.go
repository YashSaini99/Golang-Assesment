package tests

import (
	"testing"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	db := setupTestDB()
	repo := &repository.UserRepository{DB: db}
	// Hash the password before saving
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpass"), bcrypt.DefaultCost)
	user := &models.User{Username: "testuser", Password: string(hashedPassword), Role: "receptionist"}
	db.Create(user)
	// Attempt to authenticate
	token, err := service.AuthenticateUser("testuser", "testpass", repo)
	if err != nil {
		t.Errorf("Failed to authenticate user: %v", err)
	}
	if token == "" {
		t.Errorf("Expected a token, got empty string")
	}
}
