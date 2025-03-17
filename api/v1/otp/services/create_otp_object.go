package services

import (
	"strconv"

	"github.com/A-mey/GO-AUTH/api/v1/otp/constants"
	helpers "github.com/A-mey/GO-AUTH/common/helper"
)

func (os *OtpService) CreateOtpObject(emailId string) (string, error) {
	otp := helpers.RandomNumber(constants.OtpLength)
	otpString := strconv.Itoa(otp)
	fullHash, error := os.GetOtpFullHash(emailId, otpString)
	if error != nil {
		return "", error
	}
	return fullHash, nil
}
