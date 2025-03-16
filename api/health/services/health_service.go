package services

import "github.com/A-mey/GO-AUTH/api/health/interfaces"

var _ interfaces.HealthService = (*DefaultHealthService)(nil)

type DefaultHealthService struct{}
