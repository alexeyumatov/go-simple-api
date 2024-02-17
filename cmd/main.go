package main

import (
	"github.com/alexeyumatov/gin-test-api/internal/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/students", routers.GetAllStudents)
	router.Run("localhost:8000")
}
