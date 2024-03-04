package routes

import (
	"github.com/gin-gonic/gin"

)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()

	//links
	server.POST("/url/shrink", addLink)
	server.GET("/url", getAllLinks)
	server.GET("/:shorturl", getSingleUrl)
	server.DELETE("/url/:shortId", deleteUrl)
	//users
	server.POST("/register", register)
	server.POST("/register/verify",verifyEmail)
	server.POST("/login", login)

	return server
}
