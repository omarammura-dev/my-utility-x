package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"myutilityx.com/mail"
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
	server.POST("/login", login)
	server.POST("/send/email", func(ctx *gin.Context) {
		result, err := mail.SendSimpleMessage()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "oops" + err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "sent: " + result})

	})
	return server
}
