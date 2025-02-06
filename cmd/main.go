package main

import (
	"context"
	"database/sql"
	"log"
	stdhttp "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/tomygp97/weather-notifier/config"
	userHttp "github.com/tomygp97/weather-notifier/internal/delivery/http"
	weatherHttp "github.com/tomygp97/weather-notifier/internal/delivery/http"
	"github.com/tomygp97/weather-notifier/internal/infrastructure/repository"
	"github.com/tomygp97/weather-notifier/internal/usecase"
)

func main() {
	cfg := config.LoadConfig()

	// Inicializar MySQL
	db, err := sql.Open("mysql", cfg.MysqlDSN)
	if err != nil {
		log.Fatalf("❌ Error al conectar a MySQL: %v", err)
	}
	defer db.Close()

	// Inicializar Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
	ctx := context.Background()
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Error al conectar a Redis: %v", err)
	}
	log.Println("✅ Redis conectado correctamente")

	// Inicializar capas
	userRepo := repository.MySQLUserRepo{DB: db}
	userUsecase := usecase.NewUserUsecase(&userRepo)
	userHandler := userHttp.NewUserHandler(userUsecase)

	weatherRepo := repository.NewWeatherRepository()
	weatherUsecase := usecase.NewWeatherUsecase(weatherRepo, redisClient)

	weatherHandler := weatherHttp.NewWeatherHandler(weatherUsecase)

	// Configurar router
	// user
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetSingleUser).Methods("GET")
	router.HandleFunc("/users/update/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/delete/{id}", userHandler.DeleteUser).Methods("DELETE")

	// weather
	router.HandleFunc("/weather/{cityID}", weatherHandler.GetWeather).Methods("GET")
	router.HandleFunc("/waves/{cityID}", weatherHandler.GetWaves).Methods("GET")

	// Iniciar servidor
	log.Println("🚀 Servidor corriendo en el puerto", cfg.Port)
	log.Fatal(stdhttp.ListenAndServe(":"+cfg.Port, router))
}
