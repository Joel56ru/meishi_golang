package repository

import (
	"github.com/jmoiron/sqlx"
)

//Repository Struct repo
type Repository struct {
}

//VeryCuteRepository Constructor repo
func VeryCuteRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
