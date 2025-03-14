package services

import (
	"github.com/A-mey/Auth_db/api/health/constants"
)

func (s *DefaultHealthService) HealthCheck() string {
	const message string = constants.HEALTHCHECK
	return message
}
