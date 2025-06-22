package tests

import (
	"testing"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"
)

func TestUpdatePatient(t *testing.T) {
	db := setupTestDB()
	repo := &repository.PatientRepository{DB: db}
	svc := &service.PatientService{Repo: repo}
	patient := &models.Patient{Name: "Test", Age: 30, Address: "Test St", Condition: "Healthy"}
	svc.CreatePatient(patient)
	patient.Name = "Updated Name"
	err := svc.UpdatePatient(patient)
	if err != nil {
		t.Errorf("Failed to update patient: %v", err)
	}
	var updated models.Patient
	db.First(&updated, patient.ID)
	if updated.Name != "Updated Name" {
		t.Errorf("Expected name to be 'Updated Name', got '%s'", updated.Name)
	}
}
