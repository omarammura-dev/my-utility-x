package routes

import (
	"github.com/gin-gonic/gin"
	"myutilityx.com/db"
)



func RegisterRoutes() *gin.Engine{
	server := gin.Default()
	server.POST("/url/create", db.Addl)
	return server
}