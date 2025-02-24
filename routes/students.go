package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getStudents(context *gin.Context) {
	students, err := models.GetAllStudents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch students."})
		return
	}

	context.JSON(http.StatusOK, students)
}

func createStudent(context *gin.Context) {
	var student models.Student

	err := context.ShouldBindJSON(&student)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data: " + err.Error()})
		return
	}

	err = student.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save student: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Student created successfully", "student": student})
}