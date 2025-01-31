package main

import (
	"database/sql"
	"log"
	stdhttp "net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/tomygp97/weather-notifier/config"
	"github.com/tomygp97/weather-notifier/internal/delivery/http"
	"github.com/tomygp97/weather-notifier/internal/infrastructure/repository"
	"github.com/tomygp97/weather-notifier/internal/usecase"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("mysql", cfg.MysqlDSN)
	if err != nil {
		log.Fatalf("Error al conectar a MySQL: %v", err)
	}
	defer db.Close()

	// Inicializar capas
	userRepo := repository.MySQLUserRepo{DB: db}
	userUsecase := usecase.NewUserUsecase(&userRepo)
	userHandler := http.NewUserHandler(userUsecase)

	// Configurar router
	router := mux.NewRouter()
	router.HandleFunc("/users", userHandler.RegisterUser).Methods("POST")
	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", userHandler.GetSingleUser).Methods("GET")
	router.HandleFunc("/users/update/{id}", userHandler.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/delete/{id}", userHandler.DeleteUser).Methods("DELETE")

	// Iniciar servidor
	log.Println("🚀 Servidor corriendo en el puerto", cfg.Port)
	log.Fatal(stdhttp.ListenAndServe(":"+cfg.Port, router))
}
