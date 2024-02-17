package database

import "github.com/alexeyumatov/gin-test-api/internal/models"

// Mock (will be changed in the future)

var Students = []models.Student{
	{ID: "1", Name: "Alexey", Grade: 10, Rating: 4.98},
	{ID: "2", Name: "Andrey", Grade: 7, Rating: 3.12},
}
