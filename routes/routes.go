package routes

import (
	"github.com/gin-gonic/gin"
  h "github.com/lgukasyan/passplanet/handlers"
)

func Routes() *gin.Engine {
	r := gin.Default()
  r.GET("/ping", h.Ping)
  
	return r
}