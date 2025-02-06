package infrastructure

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Cambiar si es necesario
		Password: "",               // Si tiene contraseña, agregar aquí
		DB:       0,                // Usar la DB 0
	})

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error conectando a Redis: %v", err)
	}

	log.Println("✅ Redis conectado correctamente")
}

// Guardar en Redis con expiración
func SetCache(key string, value string, duration time.Duration) error {
	ctx := context.Background()
	return RedisClient.Set(ctx, key, value, duration).Err()
}

// Obtener de Redis
func GetCache(key string) (string, error) {
	ctx := context.Background()
	return RedisClient.Get(ctx, key).Result()
}
