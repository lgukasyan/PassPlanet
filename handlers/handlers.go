package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusAccepted, &requestUserBody)
}
