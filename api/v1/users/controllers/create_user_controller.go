package controllers

import (
	"net/http"

	"github.com/A-mey/GO-AUTH/api/v1/users/models"
	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) CreateNewUserController(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ctrl.UserService.CreateNewUserService(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
