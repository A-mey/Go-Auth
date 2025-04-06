package interfaces

import "github.com/gin-gonic/gin"

type OtpControllerInterface interface {
	SendOtp(c *gin.Context)
	ValidateOtp(c *gin.Context)
}
