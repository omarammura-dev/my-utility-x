package routes

import (
	"github.com/gin-gonic/gin"
	"myutilityx.com/http"
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
	authenticated.DELETE("/url/:id", deleteUrl)
	server.GET("/:shorturl", getSingleUrl)
	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"API": "WORKING"})
	})
	authenticated.POST("/user/update-role", setAdminRole)
	// server.GET("/", http.PlaygroundHandler())
	server.POST("/graphql", http.GraphQLHandler())
	server.GET("/check-db-connection", checkMongoDBConnection)
	//users
	server.POST("/user/register", register)
	server.POST("/save-sms", saveSms)
	server.POST("/contact-form", contactForm)
	server.GET("/get-sms", getSms)
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
