package services

import (
	"errors"

	"github.com/A-mey/Auth_db/api/v1/users/models"
)

func (s *DefaultUserService) CreateNewUserService(user *models.User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}
	user.Id = 1
	return nil
}
