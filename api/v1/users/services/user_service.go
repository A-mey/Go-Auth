package services

import "github.com/A-mey/Auth_db/api/v1/users/interfaces"

var _ interfaces.UserService = (*DefaultUserService)(nil)

type DefaultUserService struct{}
