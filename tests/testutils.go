package tests

import (
	"github.com/YashSaini99/Golang-Assesment/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Patient{}, &models.User{})
	return db
}
