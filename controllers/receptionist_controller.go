package controllers

import (
	"net/http"
	"strconv"

	"github.com/YashSaini99/Golang-Assesment/models"
	"github.com/YashSaini99/Golang-Assesment/repository"
	"github.com/YashSaini99/Golang-Assesment/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreatePatient godoc
// @Summary Create patient
// @Description Register a new patient (Receptionist only)
// @Tags patients
// @Accept  json
// @Produce  json
// @Param patient body models.Patient true "Patient data"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /receptionist/patients [post]
// @Security BearerAuth
func CreatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient models.Patient
		if err := c.ShouldBindJSON(&patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		repo := &repository.PatientRepository{DB: db}
		svc := &service.PatientService{Repo: repo}
		if err := svc.CreatePatient(&patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create patient"})
			return
		}
		c.JSON(http.StatusCreated, patient)
	}
}

// ListPatients godoc
// @Summary List patients
// @Description Get all patients (Receptionist/Doctor)
// @Tags patients
// @Produce  json
// @Success 200 {array} models.Patient
// @Failure 500 {object} map[string]string
// @Router /receptionist/patients [get]
// @Security BearerAuth
func ListPatients(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		repo := &repository.PatientRepository{DB: db}
		svc := &service.PatientService{Repo: repo}
		patients, err := svc.GetAllPatients()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
			return
		}
		c.JSON(http.StatusOK, patients)
	}
}

// UpdatePatient godoc
// @Summary Update patient
// @Description Update patient details (Receptionist/Doctor)
// @Tags patients
// @Accept  json
// @Produce  json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Patient data"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /receptionist/patients/{id} [put]
// @Security BearerAuth
func UpdatePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		repo := &repository.PatientRepository{DB: db}
		svc := &service.PatientService{Repo: repo}
		patient, err := svc.GetPatientByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
			return
		}
		if err := c.ShouldBindJSON(patient); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := svc.UpdatePatient(patient); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
			return
		}
		c.JSON(http.StatusOK, patient)
	}
}

// DeletePatient godoc
// @Summary Delete patient
// @Description Delete a patient (Receptionist only)
// @Tags patients
// @Param id path int true "Patient ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /receptionist/patients/{id} [delete]
// @Security BearerAuth
func DeletePatient(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		repo := &repository.PatientRepository{DB: db}
		svc := &service.PatientService{Repo: repo}
		if err := svc.DeletePatient(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete patient"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Patient deleted"})
	}
}
