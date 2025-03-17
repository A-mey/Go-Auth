package controllers

import (
	"github.com/A-mey/GO-AUTH/api/v1/otp/interfaces"
)

var _ interfaces.OtpControllerInterface = (*OtpController)(nil)

type OtpController struct {
	otpService interfaces.OtpServiceInterface
}

func NewOtpController(service interfaces.OtpServiceInterface) *OtpController {
	return &OtpController{otpService: service}
}
