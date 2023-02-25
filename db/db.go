package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB() *pgx.Conn {
	var uri string = os.Getenv("POSTGRESQL_URI")
	if uri == "" {
		log.Fatalf("unable to connect to the database, missing database (URI).")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URI"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	defer conn.Close(context.Background())

	log.Println("PostgreSQL is connected!")
	return conn
}