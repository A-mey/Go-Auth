package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *DefaultHealthController) HealthCheck(c *gin.Context) {
	message := h.HealthService.HealthCheck()
	c.JSON(http.StatusOK, gin.H{"message": message})
}
