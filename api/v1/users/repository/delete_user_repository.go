package userRepository

import (
	"errors"
	"time"

	"github.com/A-mey/GO-AUTH/api/v1/users/models"
)

var user = []models.User{
	{Id: 1, Title: "Mr", Email: "john@example.com", FirstName: "John", LastName: "Doe", Gender: 1, Dob: time.Now(), IsDeleted: false},
	{Id: 2, Title: "Ms", Email: "jane@example.com", FirstName: "Jane", LastName: "Doe", Gender: 0, Dob: time.Now(), IsDeleted: false},
}

func DeleteUserDao(id int) error {
	// Loop through the users and find the user by ID
	for i, user := range user {
		if user.Id == id {
			// Mark the user as deleted
			users[i].IsDeleted = true
			return nil
		}
	}
	return errors.New("user not found")
}
