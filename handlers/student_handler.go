package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"StudentProject/database"
	"StudentProject/models"
)

func Create(c *gin.Context) {
	var stu models.Student
	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&stu)
	c.JSON(http.StatusCreated, stu)
}

func GetAll(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var stu models.Student
	if err := database.DB.First(&stu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	c.JSON(http.StatusOK, stu)
}

func Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var stu models.Student
	if err := database.DB.First(&stu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&stu)
	c.JSON(http.StatusOK, stu)
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	database.DB.Delete(&models.Student{}, id)
	c.Status(http.StatusNoContent)
}
