// database/database.go
package database

import (
	"cinedle-backend/internal/config"
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	dbPool *pgxpool.Pool
	once   sync.Once
	ctx    context.Context
)

func connect() {
	cfg := config.LoadConfig()
	ctx = context.Background()

	var err error
	dbPool, err = pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("✅ Pool de conexões do banco de dados inicializado com sucesso!")
}

func GetDBPool() *pgxpool.Pool {
	once.Do(connect)
	return dbPool
}

func GetCtx() context.Context {
	once.Do(connect)
	return ctx
}

func CloseDBPool() {
	once.Do(connect)
	dbPool.Close()
	log.Println("✅ Pool de conexões do banco de dados fechado com sucesso!")
}
