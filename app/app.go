package app

import (
	"github.com/jmoiron/sqlx"
	"log"
	"notification-parser/config"
	"notification-parser/repository"
	"os"
	"time"
)

const (
	// When reconnecting to the server after connection failure
	reconnectDelay = 5 * time.Second
)

type App struct {
	db *sqlx.DB
}

func (a *App) Start() {
	log.Println("Hello world")
}

func (a *App) init() {
	dbConf := config.DatabaseConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DbName:   os.Getenv("DATABASE_NAME"),
		SslMode:  os.Getenv("DATABASE_SSL_MODE"),
	}
	for {
		repo, err := repository.New(dbConf)
		if err != nil {
			log.Println(err)
			select {
			case <-time.After(reconnectDelay):
			}
			continue
		}
		if err != nil {
			a.db = repo.Db
			break
		}
	}

}
func (a *App) connectQueues() {

}
