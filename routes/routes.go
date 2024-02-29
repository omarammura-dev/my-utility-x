package routes

import (
	"github.com/gin-gonic/gin"
)



func RegisterRoutes() *gin.Engine{
	server := gin.Default()
	server.POST("/url/create",addLink)
	server.GET("/url",getAllLinks)
	return server
}