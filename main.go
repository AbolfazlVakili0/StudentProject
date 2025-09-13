package main

import (
	"fmt"

	"StudentProject/database"
	"StudentProject/handlers"
	"StudentProject/middleware"
	"StudentProject/models"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger" 
	swaggerFiles "github.com/swaggo/files"   
	_ "StudentProject/docs"
)
// @title Student Management API
// @version 1.0
// @description This is a simple student management API with authentication
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func main() {
	database.Init()
	database.DB.AutoMigrate(&models.User{}, &models.Student{})

	fmt.Println("Database connected successfully!")
	r := gin.Default()

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Auth routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		// Student routes
		protected.POST("/students", handlers.CreateStudent)
		protected.GET("/students", handlers.GetAllStudents)
		protected.GET("/students/:id", handlers.GetStudentByID)
		protected.PUT("/students/:id", handlers.UpdateStudent)
		protected.DELETE("/students/:id", handlers.DeleteStudent)
	}

	r.Run(":8080")
}