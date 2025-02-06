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
		Port:      getEnv("PORT", "8080"),
		MysqlDSN:  getEnv("MYSQL_DSN", "root:@tcp(localhost:3306)/weather_notifier?parseTime=true"),
		RedisAddr: getEnv("REDIS_ADDR", "localhost:6379"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
