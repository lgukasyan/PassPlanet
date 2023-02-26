package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lgukasyan/passplanet/db"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func SignUp(c *gin.Context) {
	var requestUserBody struct {
		Name     string `json:"name"     binding:"required"`
		Lastname string `json:"lastname" binding:"required"`
		Email    string `json:"email"    binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.BindJSON(&requestUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	var q string = `INSERT INTO users(name, lastname, email, password) VALUES($1, $2, $3, $4);`
	_, err := db.DB.Exec(context.Background(), q, &requestUserBody.Name, &requestUserBody.Lastname, &requestUserBody.Email, &requestUserBody.Password)
	if err != nil {
		log.Fatalf(err.Error())
	}

	c.JSON(http.StatusAccepted, &requestUserBody)
}
