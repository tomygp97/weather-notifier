package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	MysqlDSN  string
	RedisAddr string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, usando valores por defecto")
	}

	return &Config{
		Port:      os.Getenv("PORT"),
		MysqlDSN:  os.Getenv("MYSQL_DSN"),
		RedisAddr: os.Getenv("REDIS_ADDR"),
	}
}
