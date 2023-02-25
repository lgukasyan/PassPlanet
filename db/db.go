package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectDB(){
	var uri string = os.Getenv("POSTGRESQL_URI")
	if uri == "" {
		log.Fatalf("unable to connect to the database, missing database (URI).")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URI"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	log.Println("PostgreSQL is connected!")
}