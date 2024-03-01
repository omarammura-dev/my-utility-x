package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	server.POST("/url/create", addLink)
	server.GET("/url", getAllLinks)
	server.GET("/:shorturl", getSingleUrl)
	server.POST("/test", func(ctx *gin.Context) {
		id, _ := shortid.Generate()
		ctx.JSON(http.StatusOK, id)

	})
	return server
}
