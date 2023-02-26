package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/passplanet/db"
	r "github.com/lgukasyan/passplanet/routes"
)

func main() {
	if err := db.OpenDB(); err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	var routes *gin.Engine = r.Routes()
	
	defer db.Close()

	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal(err)
	}
}
