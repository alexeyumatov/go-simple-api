package routers

import (
	"github.com/alexeyumatov/gin-test-api/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Students)
}

func GetStudentInfoByID(c *gin.Context) {
	student, err := database.GetStudentByID(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}
