package tests

import (
	"testing"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupPatientTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.Patient{})
	return db
}

func TestCreatePatientBasic(t *testing.T) {
	db := setupPatientTestDB()
	repo := &repository.PatientRepository{DB: db}
	svc := &service.PatientService{Repo: repo}
	patient := &models.Patient{Name: "Test", Age: 30, Address: "Test St", Condition: "Healthy"}
	err := svc.CreatePatient(patient)
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}
}
