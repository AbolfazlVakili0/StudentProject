package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"StudentProject/database"
	"StudentProject/models"
)

func CreateGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&grade)
	c.JSON(http.StatusCreated, grade)
}

func GetGradesByStudentID(c *gin.Context) {
	studentID, _ := strconv.Atoi(c.Param("id"))
	var grades []models.Grade
	database.DB.Where("student_id = ?", studentID).Find(&grades)
	c.JSON(http.StatusOK, grades)
}