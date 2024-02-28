package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func addLink(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{"message":"Hello,World!"})
	
}