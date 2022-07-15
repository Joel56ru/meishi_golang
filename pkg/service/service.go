package service

import (
	"meishi_golang/pkg/repository"
	"meishi_golang/senti"
)

//Authorization Интерфейс авторизации
type Authorization interface {
	ParseToken(accessToken, ip string) (senti.UserJWT, error)
	CreateUser(user senti.UserRegister) (int64, error)
}

//Location Интерфейс работы с гео
type Location interface {
	GetGeoCode(string) ([]senti.GeoCode, error)
}

//Service struct
type Service struct {
	Authorization
	Location
}

//VeryCuteService Construct service
func VeryCuteService(r *repository.Repository) *Service {
	return &Service{
		Authorization: MyAuthService(r.Authorization),
		Location:      MyLocationService(r.Location),
	}
}
