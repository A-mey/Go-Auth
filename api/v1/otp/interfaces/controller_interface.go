package interfaces

import "github.com/gin-gonic/gin"

type OtpControllerInterface interface {
	SendRegistrationOtp(c *gin.Context)
	SendResetPasswordOtp(c *gin.Context)
	ValidateOtp(c *gin.Context)
}
