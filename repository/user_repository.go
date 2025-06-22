package repository

import (
	"github.com/YashSaini99/Golang-Assesment/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
func (r *UserRepository) Create(user *models.User) error {
	return r.DB.Create(user).Error
}
