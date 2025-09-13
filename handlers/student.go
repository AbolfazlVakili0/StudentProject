package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"StudentProject/database"
	"StudentProject/models"
)

// CreateStudent godoc
// @Summary Create a new student
// @Description Create a new student with required information
// @Tags students
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <Add access token here>)
// @Param student body models.StudentRequest true "Student info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /students [post]
func CreateStudent(c *gin.Context) {
	var input models.StudentRequest
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := models.Student{
		Name:  input.Name,
		Age:   input.Age,
		Email: input.Email,
		Major: input.Major,
	}

	if err := database.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create student"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "student created successfully",
		"student": student,
	})
}

// GetAllStudents godoc
// @Summary Get all students
// @Description Get a list of all students
// @Tags students
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <Add access token here>)
// @Success 200 {array} models.Student
// @Failure 401 {object} map[string]interface{}
// @Router /students [get]
func GetAllStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

// GetStudentByID godoc
// @Summary Get student by ID
// @Description Get student information by student ID
// @Tags students
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <Add access token here>)
// @Param id path int true "Student ID"
// @Success 200 {object} models.Student
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /students/{id} [get]
func GetStudentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudent godoc
// @Summary Update student
// @Description Update student information by student ID
// @Tags students
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <Add access token here>)
// @Param id path int true "Student ID"
// @Param student body models.UpdateStudentRequest true "Student info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /students/{id} [put]
func UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	var student models.Student
	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "student not found"})
		return
	}

	var input models.UpdateStudentRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Name != nil {
		student.Name = *input.Name
	}
	if input.Age != nil {
		student.Age = *input.Age
	}
	if input.Email != nil {
		student.Email = *input.Email
	}
	if input.Major != nil {
		student.Major = *input.Major
	}

	if err := database.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "student updated successfully",
		"student": student,
	})
}

// DeleteStudent godoc
// @Summary Delete student
// @Description Delete student by student ID
// @Tags students
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token" default(Bearer <Add access token here>)
// @Param id path int true "Student ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid student id"})
		return
	}

	if err := database.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "student deleted successfully"})
}