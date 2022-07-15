package repository

import (
	"github.com/jmoiron/sqlx"
	"meishi_golang/senti"
)

//LocationPgsql имплементирует структуру для работы с Postgres
type LocationPgsql struct {
	db *sqlx.DB
}

//MyLocationPgsql Репозиторий для работы с базой
func MyLocationPgsql(db *sqlx.DB) *LocationPgsql {
	return &LocationPgsql{db: db}
}

//GetMetro Список метро
func (r *LocationPgsql) GetMetro() []senti.Metro {
	var metro []senti.Metro
	err := r.db.Select(&metro, `
	select
		metro_id, geo_lon, geo_lat, title, line_name, city, line_id
	from metro;
	`)
	if err != nil {
		return metro
	}
	return metro
}
