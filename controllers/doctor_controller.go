package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DoctorListPatients godoc
// @Summary Doctor list patients
// @Description Doctor can view all patients
// @Tags doctor
// @Produce  json
// @Success 200 {array} models.Patient
// @Failure 500 {object} map[string]string
// @Router /doctor/patients [get]
// @Security BearerAuth
func DoctorListPatients(db *gorm.DB) gin.HandlerFunc {
	return ListPatients(db)
}

// DoctorUpdatePatient godoc
// @Summary Doctor update patient
// @Description Doctor can update patient details
// @Tags doctor
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Patient data"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /doctor/patients/{id} [put]
// @Security BearerAuth
func DoctorUpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return UpdatePatient(db)
}
