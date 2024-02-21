package main

import (
	"github.com/alexeyumatov/gin-test-api/internal/models"
	"github.com/alexeyumatov/gin-test-api/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/students", routers.GetAllStudents)
	router.GET("/students/:id", routers.GetStudentInfoByID)
	router.POST("/students", routers.CreateStudent)
	router.PATCH("/students/:id", routers.UpdateStudent)
	router.DELETE("/students/:id", routers.DeleteStudent)

	router.Run("localhost:8000")
}
