package main

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	// Configurar Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Dejar vacío si no tiene contraseña
		DB:       0,  // Base de datos 0
	})

	// Probar conexión
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("❌ No se pudo conectar a Redis: %v", err)
	}
	fmt.Println("✅ Conectado a Redis!")

	// Guardar un valor en Redis
	err = rdb.Set(ctx, "test_key", "Hola Redis desde Go!", 0).Err()
	if err != nil {
		log.Fatalf("❌ Error al guardar en Redis: %v", err)
	}

	// Obtener el valor de Redis
	val, err := rdb.Get(ctx, "test_key").Result()
	if err != nil {
		log.Fatalf("❌ Error al obtener el valor: %v", err)
	}

	fmt.Println("🔹 Valor obtenido de Redis:", val)
}
