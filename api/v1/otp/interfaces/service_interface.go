package interfaces

type OtpServiceInterface interface {
	CreateOtpObject(emailId string) (string, error)
	ValidateOtp(emailId string, hash, otp string) (bool, error)
	GetOtpFullHash(emailId, otp string) (string, error)
}
