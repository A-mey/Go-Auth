package interfaces

import "github.com/A-mey/Auth_db/api/v1/users/models"

type UserService interface {
	CreateNewUserService(user *models.User) error
	DeleteUserService(Id int) error
}

type DefaultUserService struct{}
