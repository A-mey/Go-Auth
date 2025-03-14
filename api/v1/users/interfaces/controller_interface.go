package interfaces

import "github.com/gin-gonic/gin"

type UserControllerInterface interface {
	CreateNewUserController(c *gin.Context)
	DeleteUserController(c *gin.Context)
}
