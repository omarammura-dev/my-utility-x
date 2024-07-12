package routes

import (
	"net/http"

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
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"API": "WORKING"})
	})
	server.GET("/check-db-connection", checkMongoDBConnection)
	//users
	server.POST("/user/register", register)
	server.POST("/save-sms", saveSms)
	server.GET("/user/verify", verifyEmail)
	server.POST("/user/login", login)
	server.POST("/user/reset-password", resetPassword)
	server.POST("/user/reset-password/confirm", resetPasswordVerify)

	//expenses statistics
	authenticated.POST("/expense/add", addExpense)
	authenticated.PUT("/expense/update/:id", updateExpense)
	authenticated.GET("/expense/all", GetAllExpenses)
	return server
}
