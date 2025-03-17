package services

import "github.com/A-mey/GO-AUTH/api/v1/otp/interfaces"

var _ interfaces.OtpServiceInterface = (*OtpService)(nil)

type OtpService struct {
}
