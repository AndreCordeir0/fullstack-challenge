package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Password string
	Host     string
	Database string
}

func GetConnection() *sql.DB {
	config := &Config{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
	}
	connectionStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.Username, config.Password, config.Host, config.Database)
	log.Default().Println("Connection", connectionStr)
	sqlDb, err := sql.Open("postgres", connectionStr)

	if err != nil {
		log.Fatal("Error connecting in database", err)
	}
	log.Default().Println("Connected in database")
	return sqlDb
}
