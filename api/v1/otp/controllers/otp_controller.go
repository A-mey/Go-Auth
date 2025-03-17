package controllers

import (
	otpInterfaces "github.com/A-mey/GO-AUTH/api/v1/otp/interfaces"
)

var _ otpInterfaces.OtpControllerInterface = (*OtpController)(nil)

type OtpController struct {
	otpService otpInterfaces.OtpServiceInterface
}

func NewOtpController(service otpInterfaces.OtpServiceInterface) *OtpController {
	return &OtpController{otpService: service}
}
