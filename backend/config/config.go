package config

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)


func LoadEnv() {
	_ = godotenv.Load(".env")
	_ = godotenv.Load(filepath.Join("..", ".env"))
}


func PostgresDSN() (string, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if user == "" || password == "" || host == "" || port == "" || dbname == "" || sslmode == "" {
		return "", fmt.Errorf("database env incomplete: set DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME, DB_SSLMODE (e.g. via backend/.env)")
	}

	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, password),
		Host:   host + ":" + port,
		Path:   "/" + dbname,
	}
	q := u.Query()
	q.Set("sslmode", sslmode)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
