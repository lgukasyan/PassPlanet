package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/passplanet/db"
)

func main() {
	r := gin.Default()

	db.ConnectDB()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
