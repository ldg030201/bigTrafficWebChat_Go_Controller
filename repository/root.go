package repository

import (
	"chat_controller_server/config"
	"chat_controller_server/repository/kafka"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	cfg *config.Config

	db *sql.DB

	Kafka *kafka.Kafka
}

const (
	room       = " bigTrafficWebChat.room "
	chat       = " bigTrafficWebChat.chat "
	serverInfo = " bigTrafficWebChat.serverInfo "
)

func NewRepository(cfg *config.Config) (*Repository, error) {
	r := &Repository{cfg: cfg}
	var err error

	if r.db, err = sql.Open(cfg.DB.Database, cfg.DB.URL); err != nil {
		return nil, err
	} else if r.Kafka, err = kafka.NewKafka(cfg); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}
