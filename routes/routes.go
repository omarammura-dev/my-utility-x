package routes

import (
	"github.com/gin-gonic/gin"
	"myutilityx.com/middlewares"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	server.Use(middlewares.CORSMiddleware())
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	//links
	authenticated.POST("/url/shrink", addLink)
	authenticated.GET("/url", getAllLinks)
	authenticated.DELETE("/url/:shortId", deleteUrl)
	server.GET("/:shorturl", getSingleUrl)

	//users
	server.POST("user/register", register)
	server.POST("user/verify", verifyEmail)
	server.POST("user/login", login)
	server.POST("user/reset-password", resetPassword)
	server.POST("user/reset-password/confirm", resetPasswordVerify)
	return server
}
