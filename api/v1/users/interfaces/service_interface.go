package interfaces

import "github.com/A-mey/GO-AUTH/api/v1/users/models"

type UserService interface {
	CreateNewUserService(user *models.User) error
	DeleteUserService(Id int) error
}

type DefaultUserService struct{}
