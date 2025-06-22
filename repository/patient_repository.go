package repository

import (
	"github.com/YashSaini99/Golang-Assesment/models"
	"gorm.io/gorm"
)

type PatientRepository struct {
	DB *gorm.DB
}

func (r *PatientRepository) Create(patient *models.Patient) error {
	return r.DB.Create(patient).Error
}

func (r *PatientRepository) FindAll() ([]models.Patient, error) {
	var patients []models.Patient
	err := r.DB.Find(&patients).Error
	return patients, err
}

func (r *PatientRepository) FindByID(id uint) (*models.Patient, error) {
	var patient models.Patient
	err := r.DB.First(&patient, id).Error
	return &patient, err
}

func (r *PatientRepository) Update(patient *models.Patient) error {
	return r.DB.Save(patient).Error
}

func (r *PatientRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Patient{}, id).Error
}
