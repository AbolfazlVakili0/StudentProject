package main

import (
    "fmt"
	"StudentProject/database"
	"StudentProject/handlers"
	"StudentProject/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
    err := database.DB.AutoMigrate(&models.Student{})
	if err != nil {
		panic("migration failed: " + err.Error())
	}
	
	fmt.Println("Database connected successfully!")
	r := gin.Default()

	r.POST("/students", handlers.Create)
	r.GET("/students", handlers.GetAll)
	r.GET("/students/:id", handlers.GetByID)
	r.PUT("/students/:id", handlers.Update)
	r.DELETE("/students/:id", handlers.Delete)

	r.Run(":8080")
}