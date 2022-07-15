package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

//Config struct connection db
type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

//CuteConnectionDB connect to database
func CuteConnectionDB(k Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(`host=%s port=%s dbname=%s user=%s password=%s sslmode=%s`, k.Host, k.Port, k.DBName, k.Username, k.Password, k.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}
