package config

import (
	"fmt"
	"os"
	"sync"

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

var (
	config *Config
	once   sync.Once
)

// LoadConfig inicializa a configuração apenas uma vez
func LoadConfig() *Config {
	once.Do(func() {
		// Tenta carregar .env (se existir)
		if err := godotenv.Load(); err != nil {
			fmt.Println("⚠️ Não foi possível carregar .env ou arquivo não existe. Usando variáveis de ambiente do sistema.")
		}

		config = &Config{
			DatabaseURL: os.Getenv("DB_URL"),
			Port:        os.Getenv("PORT"),
			DBUser:      os.Getenv("DB_USER"),
			DBPassword:  os.Getenv("DB_PASSWORD"),
			DBName:      os.Getenv("DB_NAME"),
			DBHost:      os.Getenv("DB_HOST"),
			DBPort:      os.Getenv("DB_PORT"),
		}
	})

	return config
}
