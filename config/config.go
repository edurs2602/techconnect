// config/config.go
package config

import (
	"log"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

func Load() Config {
	return Config{
		DatabaseURL: mustGet("DATABASE_URL"),
		Port:        getOrDefault("PORT", "8080"),
		JWTSecret:   mustGet("JWT_SECRET"),
	}
}

func NewDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalf("erro ao conectar no banco: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("banco inacessível: %v", err)
	}
	return db
}

func mustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("variável de ambiente obrigatória ausente: %s", key)
	}
	return v
}

func getOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
