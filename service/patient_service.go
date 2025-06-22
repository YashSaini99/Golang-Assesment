package service

import (
	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
)

type PatientService struct {
	Repo *repository.PatientRepository
}

func (s *PatientService) CreatePatient(patient *models.Patient) error {
	return s.Repo.Create(patient)
}

func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
	return s.Repo.FindAll()
}

func (s *PatientService) GetPatientByID(id uint) (*models.Patient, error) {
	return s.Repo.FindByID(id)
}

func (s *PatientService) UpdatePatient(patient *models.Patient) error {
	return s.Repo.Update(patient)
}

func (s *PatientService) DeletePatient(id uint) error {
	return s.Repo.Delete(id)
}
