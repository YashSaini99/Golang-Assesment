package tests

import (
	"testing"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"
)

func TestDeletePatient(t *testing.T) {
	db := setupTestDB()
	repo := &repository.PatientRepository{DB: db}
	svc := &service.PatientService{Repo: repo}
	patient := &models.Patient{Name: "Test", Age: 30, Address: "Test St", Condition: "Healthy"}
	svc.CreatePatient(patient)
	err := svc.DeletePatient(patient.ID)
	if err != nil {
		t.Errorf("Failed to delete patient: %v", err)
	}
	var deleted models.Patient
	result := db.First(&deleted, patient.ID)
	if result.Error == nil {
		t.Errorf("Expected patient to be deleted, but found in DB")
	}
}
