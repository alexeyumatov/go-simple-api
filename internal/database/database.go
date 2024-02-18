package database

import (
	"errors"
	"github.com/alexeyumatov/gin-test-api/internal/models"
)

var errorNotFound = errors.New("Object Not Found")

// Mock (will be changed in the future)
var Students = []models.Student{
	{ID: "1", Name: "Alexey", Grade: 10, Rating: 4.98},
	{ID: "2", Name: "Andrey", Grade: 7, Rating: 3.12},
}

func GetStudentByID(id string) (*models.Student, error) {
	for _, v := range Students {
		if v.ID == id {
			return &v, nil
		}
	}
	return nil, errorNotFound
}
