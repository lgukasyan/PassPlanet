package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func OpenDB() (error) {
	var uri string = os.Getenv("POSTGRESQL_URI")

	if uri == "" {
		log.Fatalf("unable to connect to the database, missing database (URI).")
	}

	var err error
	DB, err = pgx.Connect(context.Background(), uri)
	log.Println("PostgreSQL is connected!")
	return err
}

func Close() error {
	return DB.Close(context.Background())
}