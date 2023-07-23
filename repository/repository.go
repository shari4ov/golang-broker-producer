package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"notification-parser/config"
)

type Repository struct {
	Db *sqlx.DB
}

func (r *Repository) GetReminders() {
	return
}
func New(d config.DatabaseConfig) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.DbName, d.SslMode)

	db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}
