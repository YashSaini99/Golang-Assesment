package routes

import (
	"github.com/YashSaini99/Golang-Assesment/controllers"
	"github.com/YashSaini99/Golang-Assesment/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", controllers.Login(db))

	receptionist := r.Group("/receptionist")
	receptionist.Use(middleware.AuthMiddleware("receptionist"))
	{
		receptionist.POST("/patients", controllers.CreatePatient(db))
		receptionist.GET("/patients", controllers.ListPatients(db))
		receptionist.PUT("/patients/:id", controllers.UpdatePatient(db))
		receptionist.DELETE("/patients/:id", controllers.DeletePatient(db))
	}

	doctor := r.Group("/doctor")
	doctor.Use(middleware.AuthMiddleware("doctor"))
	{
		doctor.GET("/patients", controllers.DoctorListPatients(db))
		doctor.PUT("/patients/:id", controllers.DoctorUpdatePatient(db))
	}
}
