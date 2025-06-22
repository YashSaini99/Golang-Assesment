package tests

import (
	"testing"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"
)

func TestCreatePatient(t *testing.T) {
	db := setupTestDB()
	repo := &repository.PatientRepository{DB: db}
	svc := &service.PatientService{Repo: repo}
	patient := &models.Patient{Name: "Test", Age: 30, Address: "Test St", Condition: "Healthy"}
	err := svc.CreatePatient(patient)
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}
}
