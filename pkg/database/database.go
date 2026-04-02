package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {
	const MAX_RETRY_DB_CON = 5

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		getEnv("DB_HOST", "db"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "manomano"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	for i := 1; i <= MAX_RETRY_DB_CON; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("db not ready, attempt %d/%d: %v", i, MAX_RETRY_DB_CON, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
