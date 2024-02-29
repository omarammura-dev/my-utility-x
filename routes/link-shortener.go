package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"myutilityx.com/models"
)


func addLink(ctx *gin.Context){
	var link models.Link
	err := ctx.ShouldBindJSON(&link)

	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"failed to parse the link object!"})	
		return
	}

	err = link.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"failed to save the link!"})		
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"link saved success!"})
}


func getAllLinks(ctx *gin.Context){

	var link models.Link
	err,linkList := link.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"failed to get the links!"})		
	}
	
	ctx.JSON(http.StatusOK,linkList)

}