package database

import (
	"cinedle-backend/internal/config"
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	connection *pgx.Conn
	ctx        context.Context
}

func New() *DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Configuração inválida")
	}
	db, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	return &DB{
		connection: db,
		ctx:        context.Background(),
	}
}
func (db *DB) GetConnection() *pgx.Conn {
	return db.connection
}
func (db *DB) GetContext() context.Context {
	return db.ctx
}
func (db *DB) Close() error {
	return db.connection.Close(db.ctx)
}
