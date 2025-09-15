package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	DBUser      string
	DBPassword  string
	DBName      string
	DBHost      string
	DBPort      string
}

func LoadConfig() (*Config, error) {

	// Tenta carregar .env (se existir)
	if err := godotenv.Load(); err != nil {
		// Não fatal se o arquivo não existir; apenas loga
		fmt.Println("⚠️ Não foi possível carregar .env ou arquivo não existe. Usando variáveis de ambiente do sistema.")
	}
	DatabaseURL := os.Getenv("DB_URL")

	Port := os.Getenv("PORT")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	return &Config{
		DatabaseURL: DatabaseURL,
		Port:        Port,
		DBUser:      DBUser,
		DBPassword:  DBPassword,
		DBName:      DBName,
		DBHost:      DBHost,
		DBPort:      DBPort,
	}, nil
}
