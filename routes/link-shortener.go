package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"myutilityx.com/models"
)

func addLink(ctx *gin.Context) {
	link, err := models.InitLink()
	err = ctx.ShouldBindJSON(&link)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to parse the link object!"})
		return
	}

	err = link.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save the link!"})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "link saved success!"})
}

func getAllLinks(ctx *gin.Context) {

	var link models.Link
	linkList, err := link.GetAll()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get the links!"})
	}

	ctx.JSON(http.StatusOK, linkList)

}

func getSingleUrl(ctx *gin.Context) {
	shortUrl := ctx.Param("shorturl")
	if strings.HasPrefix(shortUrl, "U") {
		l, err := models.GetSingleAndIncreaseClicks(shortUrl)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
		}
	
		ctx.Redirect(http.StatusMovedPermanently,l.Url)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
	}
}
