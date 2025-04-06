package routes

import (
	"github.com/A-mey/GO-AUTH/api/v1/otp/controllers"
	"github.com/A-mey/GO-AUTH/api/v1/otp/interfaces"
	"github.com/A-mey/GO-AUTH/api/v1/otp/services"
	schemaValidationInterfaces "github.com/A-mey/GO-AUTH/api/v1/schema-validator/interfaces"
	RoutesInterface "github.com/A-mey/GO-AUTH/common/interfaces"
	"github.com/gin-gonic/gin"
)

var _ RoutesInterface.RoutesInterface = (*otpRoutes)(nil)

type otpRoutes struct {
	schemaValidatorMiddleware schemaValidationInterfaces.SchemaValidatorMiddlewareInterfaces
}

func NewOtpRoutes(middleware schemaValidationInterfaces.SchemaValidatorMiddlewareInterfaces) RoutesInterface.RoutesInterface {
	return &otpRoutes{schemaValidatorMiddleware: middleware}
}

func (hr *otpRoutes) InitializeRoutes(r *gin.Engine) {

	otpService := &services.OtpService{}

	var otpController interfaces.OtpControllerInterface = controllers.NewOtpController(otpService)
	otpRoutes := r.Group("/otp")

	r.Use(hr.schemaValidatorMiddleware.ValidationRequestMiddleware(map[string]string{"name": "John Doe", "age": "30", "email": "john@example.com"}))

	{
		otpRoutes.POST("/send", otpController.SendOtp)
		otpRoutes.POST("/validate", otpController.ValidateOtp)
	}
}
