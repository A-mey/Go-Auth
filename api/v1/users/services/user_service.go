package services

import "github.com/A-mey/GO-AUTH/api/v1/users/interfaces"

var _ interfaces.UserService = (*DefaultUserService)(nil)

type DefaultUserService struct{}
