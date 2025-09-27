// database/database.go
package database

import (
	"cinedle-backend/internal/config"
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
)

var (
	db   *pgx.Conn
	once sync.Once
	ctx  context.Context
)

func connect() {
	cfg := config.LoadConfig()
	ctx = context.Background()

	var err error
	db, err = pgx.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("✅ Conectado ao banco de dados com sucesso!")
}

func GetDB() *pgx.Conn {
	once.Do(connect)
	return db
}
func GetCtx() context.Context {
	once.Do(connect)
	return ctx
}

func CloseDB() {
	err := db.Close(ctx)
	if err != nil {
		log.Fatalf("❌ Erro ao fechar a conexão do banco de dados: %v", err)
	}
	log.Println("✅ Conexão do banco de dados fechada com sucesso!")
}
