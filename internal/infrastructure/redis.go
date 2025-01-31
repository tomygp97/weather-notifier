package infrastructure

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedis(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("⚠️ No se pudo conectar a Redis: ", err)
	} else {
		log.Println("✅ Conectado a Redis")
	}

	return rdb
}
