package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_HOST     = "127.0.0.1"
	DB_USER     = "root"
	DB_PASSWORD = "root"
	DB_NAME     = "lista_compras_db"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASSWORD,
		DB_HOST, DB_NAME)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	log.Println("Conexao com BD ok")
	return db, nil
}
