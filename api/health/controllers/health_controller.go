package controllers

import "github.com/A-mey/Auth_db/api/health/interfaces"

var _ interfaces.HealthController = (*DefaultHealthController)(nil)

type DefaultHealthController struct {
	HealthService interfaces.HealthService
}

func NewHealthController(service interfaces.HealthService) *DefaultHealthController {
	return &DefaultHealthController{HealthService: service}
}
