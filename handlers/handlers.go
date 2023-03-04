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

var (
	err   error
	q     string
	valid bool
	rows  pgx.Rows
	row   pgx.Row
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func SignUp(c *gin.Context) {
	var requestUserBody struct {
		Name     string `json:"name"     binding:"required"`
		Lastname string `json:"lastname" binding:"required"`
		Email    string `json:"email"    binding:"required"`
		Key      string `json:"key"      binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err = c.BindJSON(&requestUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	q = `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1);`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.Email)
	err = row.Scan(&valid)

	if valid || err != nil {
		log.Println("Email already exists")
		return
	}

	err = u.HashPassword(&requestUserBody.Password)
	if err != nil {
		log.Fatalf("error hashing the password %s", err.Error())
	}

	q = `INSERT INTO users(name, lastname,  email, password, key) VALUES($1, $2, $3, $4, $5);`
	_, err = db.DB.Exec(context.Background(), q,
		&requestUserBody.Name,
		&requestUserBody.Lastname,
		&requestUserBody.Email,
		&requestUserBody.Password,
		&requestUserBody.Key,
	)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	c.JSON(http.StatusAccepted, &requestUserBody)
}

func SignIn(c *gin.Context) {
	var requestUserBody struct {
		Email    string `json:"email"    binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err = c.BindJSON(&requestUserBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	var password string

	q = `SELECT (password) FROM users WHERE email=$1;`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.Email)
	err = row.Scan(&password)

	if err == pgx.ErrNoRows {
		log.Println("User not found")
		return
	}

	err = u.ComparePassword(&password, &requestUserBody.Password)
	if err != nil {
		log.Println("error, incorrect password")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "logged",
	})
}

func CreateNewPassword(c *gin.Context) {
	var requestUserBody struct {
		User_id     int    `json:"user_id"`
		Url         string `json:"url"`
		IB64        string `json:"icon_base64data"`
		Title       string `json:"title"   					binding:"required"`
		Description string `json:"description" 			binding:"required"`
		Password    string `json:"password" 				binding:"required"`
	}

	var err error

	if err = c.BindJSON(&requestUserBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	var (
		q     string
		row   pgx.Row
		valid bool
	)

	q = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1)`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.User_id)
	err = row.Scan(&valid)

	if !valid || err != nil {
		log.Println("User not found")
		return
	}

	q = `INSERT INTO passwords(user_id, title, description, password) VALUES($1, $2, $3, $4);`
	_, err = db.DB.Exec(context.Background(), q,
		&requestUserBody.User_id,
		&requestUserBody.Title,
		&requestUserBody.Description,
		&requestUserBody.Password,
	)

	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "created",
	})
}

func DeletePassword(c *gin.Context) {
	var requestUserBody struct {
		User_id     int `json:"user_id" binding:"required"`
		Password_id int `json:"password_id" binding:"required"`
	}

	if err = c.BindJSON(&requestUserBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	q = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1);`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.User_id)
	err = row.Scan(&valid)

	if !valid || err != nil {
		log.Println("User doesn't exist")
		return
	}

	q = `DELETE FROM passwords WHERE password_id = $1 AND user_id = $2;`
	res, err := db.DB.Exec(context.Background(), q, &requestUserBody.Password_id, &requestUserBody.User_id)

	if err != nil {
		log.Println("error deleting the password")
		return
	}

	var count int64 = res.RowsAffected()
	if count == 0 {
		log.Println("password not found for deletion")
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "deleted",
	})
}

func GetAllPass(c *gin.Context) {
	var requestUserBody struct {
		User_id int `json:"user_id"`
	}

	if err = c.BindJSON(&requestUserBody); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error binding json"})
		return
	}

	q = `SELECT EXISTS(SELECT 1 FROM users WHERE user_id=$1);`
	row = db.DB.QueryRow(context.Background(), q, &requestUserBody.User_id)
	err = row.Scan(&valid)

	if !valid || err != nil {
		log.Println("User doesn't exist")
		return
	}

	q = `SELECT user_id, title, description, password FROM passwords WHERE user_id=$1;`
	rows, err = db.DB.Query(context.Background(), q, &requestUserBody.User_id)
	if err != nil {
		log.Println("error getting all the passwords")
		return
	}
	defer rows.Close()

	var passl *([]models.Password) = &([]models.Password{})

	for rows.Next() {
		var p *models.Password = &models.Password{}
		if err = rows.Scan(&p.User_id, &p.Title, &p.Description, &p.Password); err != nil {
			log.Println(err.Error())
			log.Println("error getting the passwords")
			return
		}
		*passl = append(*passl, *p)
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": passl,
	})
}
