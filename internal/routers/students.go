package routers

import (
	"github.com/alexeyumatov/gin-test-api/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Students)
}
