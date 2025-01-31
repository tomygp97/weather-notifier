package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Conectado a la base de datos MySQL")
	return db, nil
}
