package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func LoadConfig(envFile string) (*Config, error) {
	// Tenta carregar .env (se existir)
	if err := godotenv.Load(envFile); err != nil {
		// Não fatal se o arquivo não existir; apenas loga
		fmt.Println("⚠️ Não foi possível carregar .env ou arquivo não existe. Usando variáveis de ambiente do sistema.")
	}
	postgresURL := os.Getenv("DB_URL")
	if postgresURL == "" {
		postgresURL = "postgres://postgres:@localhost:5432/whatsapp_db?sslmode=disable"
	}

	return &Config{
		DatabaseURL: postgresURL,
	}, nil
}
