package routes

import (
	"github.com/A-mey/GO-AUTH/api/middlewares"
	"github.com/A-mey/GO-AUTH/api/v1/otp/controllers"
	"github.com/A-mey/GO-AUTH/api/v1/otp/interfaces"
	"github.com/A-mey/GO-AUTH/api/v1/otp/services"
	RoutesInterface "github.com/A-mey/GO-AUTH/common/interfaces"
	"github.com/gin-gonic/gin"
)

var _ RoutesInterface.RoutesInterface = (*otpRoutes)(nil)

type otpRoutes struct{}

func (hr *otpRoutes) InitializeRoutes(r *gin.Engine) {

	r.Use(middlewares.ValidationRequestMiddleware(map[string]string{"name": "John Doe", "age": "30", "email": "john@example.com"}))

	otpService := &services.OtpService{}

	var otpController interfaces.OtpControllerInterface = controllers.NewOtpController(otpService)
	userRoutes := r.Group("/otp")
	{
		userRoutes.POST("/registration", otpController.SendRegistrationOtp)
		userRoutes.POST("/resetPassword", otpController.SendResetPasswordOtp)
	}
}
