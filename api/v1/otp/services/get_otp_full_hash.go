package services

import (
	"time"

	"github.com/A-mey/GO-AUTH/api/v1/otp/constants"
	"github.com/A-mey/GO-AUTH/libs/encryption"
)

func (os *OtpService) GetOtpFullHash(emailId, otp string) (string, error) {
	otpValidationTime := constants.OtpValidationTime // in milliseconds
	expires := time.Now().Add(time.Duration(otpValidationTime) * time.Millisecond)
	dataToBeHashed := emailId + otp + expires.String()
	hash := encryption.HMAC(dataToBeHashed)
	fullHash := hash + "." + expires.String()
	return fullHash, nil
}
