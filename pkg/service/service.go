package service

import (
	"meishi_golang/pkg/repository"
)

//Service struct
type Service struct {
}

//VeryCuteService Construct service
func VeryCuteService(r *repository.Repository) *Service {
	return &Service{}
}
