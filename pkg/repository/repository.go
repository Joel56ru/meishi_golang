package repository

import (
	"github.com/jmoiron/sqlx"
	"meishi_golang/senti"
)

//Authorization Интерфейс авторизации
type Authorization interface {
	CreateUser(user senti.UserRegister) (int64, error) //Создание пользователя
}

//Location Интерфейс Локации
type Location interface {
	GetMetro() []senti.Metro //Список метро
}

//Repository Struct repo
type Repository struct {
	Authorization
	Location
}

//VeryCuteRepository Constructor repo
func VeryCuteRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: MyAuthPgsql(db),
		Location:      MyLocationPgsql(db),
	}
}
