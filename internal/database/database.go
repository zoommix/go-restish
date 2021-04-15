package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	maxConns = 4
	minConns = 2
)

func InitDatabase() (*pgxpool.Pool, error) {
	fmt.Println("Connection to DB...")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbName, dbPassword)

	config, err := pgxpool.ParseConfig(connStr)

	if err != nil {
		panic(err)
	}

	config.MaxConns = maxConns
	config.MinConns = minConns

	db, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		return db, err
	}

	return db, nil
}
