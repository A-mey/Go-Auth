package userRepository

import "github.com/A-mey/GO-AUTH/api/v1/users/models"

// Simulated in-memory user store
var users = []models.User{
	{Id: 1, FirstName: "John", LastName: "Doe", Email: "john@example.com"},
	{Id: 2, FirstName: "Jane", LastName: "Doe", Email: "jane@example.com"},
}

// func CreateNewUserDao(user *models.User) error {
// 	// for i, user := range users {
// 	// 	if user.Id == id {
// 	// 		users = append(users[:i], users[i+1:]...)
// 	// 		return nil
// 	// 	}
// 	// }
// 	users = append(users, *user)
// 	// Return error if user is not found
// 	// return errors.New("user not found")
// 	return nil
// }

func CreateUser(user *models.User) error {
	// result := sql.DB.Create(user)
	return nil
}
