package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/lgukasyan/passplanet/db"
	"github.com/lgukasyan/passplanet/models"
	_ "github.com/lgukasyan/passplanet/models"
	u "github.com/lgukasyan/passplanet/utils"
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

	var err error

	if err = c.BindJSON(&requestUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	var q string
	var row pgx.Row
	var user *models.User = &models.User{}

	q = `SELECT (email) FROM users WHERE email=$1;`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.Email)
	err = row.Scan(&user.Email)

	if err != pgx.ErrNoRows {
		log.Println("email exists")
		return
	}

	err = u.HashPassword(&requestUserBody.Password)
	
	if err != nil {
		log.Fatalf("error hashing the password %s", err.Error())
	}

	q = `INSERT INTO users(name, lastname, email, password) VALUES($1, $2, $3, $4);`
	_, err = db.DB.Exec(context.Background(), q, &requestUserBody.Name, &requestUserBody.Lastname, &requestUserBody.Email, &requestUserBody.Password)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	c.JSON(http.StatusAccepted, &requestUserBody)
}
