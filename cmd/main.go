// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"log"
	"os"

	"github.com/YashSaini99/Golang-Assesment/config"
	"github.com/YashSaini99/Golang-Assesment/controllers"
	_ "github.com/YashSaini99/Golang-Assesment/docs"
	"github.com/YashSaini99/Golang-Assesment/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on environment variables.")
	}
	db := config.SetupDatabase()
	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.POST("/register", controllers.Register(db))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Golang Assessment API! Visit /swagger/index.html for API docs."})
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
