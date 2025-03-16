package services

import (
	"github.com/A-mey/GO-AUTH/api/health/constants"
)

func (s *DefaultHealthService) HealthCheck() string {
	const message string = constants.HEALTHCHECK
	return message
}
