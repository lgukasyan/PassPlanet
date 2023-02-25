package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/lgukasyan/passplanet/db"
	r "github.com/lgukasyan/passplanet/routes"
)

type Application struct {
	DB *pgx.Conn
}

func main() {
	var app Application
	app.DB = db.ConnectDB()
	
	var routes *gin.Engine = r.Routes()

	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal(err)
	}
}
