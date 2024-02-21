package routers

import (
	"net/http"

	"github.com/alexeyumatov/gin-test-api/internal/models"
	"github.com/alexeyumatov/gin-test-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func findStudent(id string) (models.Student, error) {
	var student models.Student
	if err := models.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return student, err
	}
	return student, nil
}

func GetAllStudents(c *gin.Context) {
	var students []models.Student
	models.DB.Find(&students)
	c.IndentedJSON(http.StatusOK, students)
}

func GetStudentInfoByID(c *gin.Context) {
	student, err := findStudent(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var newStudent schemas.CreateStudentSchema

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	student := models.Student{
		Name:   newStudent.Name,
		Grade:  newStudent.Grade,
		Rating: newStudent.Rating,
	}
	models.DB.Create(&student)
	c.IndentedJSON(http.StatusCreated, student)
}

func UpdateStudent(c *gin.Context) {
	student, err := findStudent(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	var updatedStudent schemas.UpdateStudentSchema

	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	models.DB.Model(student).Updates(updatedStudent)

	c.IndentedJSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	student, err := findStudent(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	models.DB.Delete(&student)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}
