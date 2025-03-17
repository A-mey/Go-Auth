package services

import (
	"errors"
	"strings"
	"time"

	"github.com/A-mey/GO-AUTH/libs/encryption"
)

func (os *OtpService) ValidateOtp(emailId string, hash, otp string) (bool, error) {
	hashValueAndExpires := strings.Split(hash, ".")
	if len(hashValueAndExpires) != 2 {
		return false, errors.New("invalid hash format")
	}
	hashValue, expires := hashValueAndExpires[0], hashValueAndExpires[1]
	timeNow := time.Now()
	expiresTime, err := time.Parse(time.RFC3339, expires)
	if err != nil {
		return false, err
	}
	if timeNow.After(expiresTime) {
		return false, nil
	}
	dataToBeHashed := emailId + otp + expires
	newHash := encryption.HMAC(dataToBeHashed)
	if newHash == hashValue {
		return true, nil
	}
	return false, nil
}
