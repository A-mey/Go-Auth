package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctrl *UserController) DeleteUserController(c *gin.Context) {
	// Get user ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Call the service layer
	if err := ctrl.UserService.DeleteUserService(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
