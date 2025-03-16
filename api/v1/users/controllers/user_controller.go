package controllers

import "github.com/A-mey/GO-AUTH/api/v1/users/interfaces"

var _ interfaces.UserControllerInterface = (*UserController)(nil)

type UserController struct {
	UserService interfaces.UserService
}

// Constructor for UserController
func NewUserController(service interfaces.UserService) *UserController {
	return &UserController{UserService: service}
}
