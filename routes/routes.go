package routes

import (
	"github.com/gin-gonic/gin"
	h "github.com/lgukasyan/passplanet/handlers"
)

func Routes() *gin.Engine {
	
	r := gin.Default()
	
	r.GET("/ping", h.Ping)
	r.POST("/sign-up", h.SignUp)
	r.POST("/sign-in", h.SignIn)
	r.POST("/create", h.CreateNewPassword)
	r.POST("/delete", h.DeletePassword)
	r.POST("/get-all", h.GetAllPass)

	return r
}
