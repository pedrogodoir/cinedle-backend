package database

import (
	"cinedle-backend/internal/config"
	"context"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	connection *gorm.DB
	ctx        context.Context
}

func New() *DB {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic("Configuração inválida")
	}
	db, err := gorm.Open(postgres.Open("host="+cfg.DBHost+" user="+
		cfg.DBUser+" password="+cfg.DBPassword+" dbname="+
		cfg.DBName+" port="+cfg.DBPort+" sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	return &DB{
		connection: db,
		ctx:        context.Background(),
	}
}
func (db *DB) GetConnection() *gorm.DB {
	return db.connection
}
func (db *DB) GetContext() context.Context {
	return db.ctx
}
func (db *DB) Close() error {
	sqlDB, err := db.connection.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
