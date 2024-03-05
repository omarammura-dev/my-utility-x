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
	server.POST("user/register", register)
	server.POST("user/verify",verifyEmail)
	server.POST("user/login", login)
	server.POST("user/reset-password",resetPassword)
	server.POST("user/reset-password/verify",resetPasswordVerify)
	return server
}
