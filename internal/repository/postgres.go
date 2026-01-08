package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB() *sql.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dns := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbName)

	db, err := sql.Open("pgx", dns)
	if err != nil {
		log.Fatalf("Failed to establish connection to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Database is unavailable: %v", err)
	}

	fmt.Println("Successfully connected to the database")
	return db
}
