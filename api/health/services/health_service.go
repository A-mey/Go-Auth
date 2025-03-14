package services

import "github.com/A-mey/Auth_db/api/health/interfaces"

var _ interfaces.HealthService = (*DefaultHealthService)(nil)

type DefaultHealthService struct{}
