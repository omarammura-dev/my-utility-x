package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myutilityx.com/models"
)

func addLink(ctx *gin.Context) {

	link, err := models.InitLink()
	if err != nil {
		log.Fatalf("Something went wrong... %v", err)
	}
	err = ctx.ShouldBindJSON(&link)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse the link object!"})
		return
	}

	userId,exist := ctx.Get("userId")
	if !exist {

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Oops,something went wrong!"})
	}
	link.UserId = userId.(primitive.ObjectID)
	err = link.Save()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to save the link!"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "link saved success!"})
}

func getAllLinks(ctx *gin.Context) {

	userId,exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(http.StatusInternalServerError,gin.H{"message":"Oops! something went wrong!"})
	}

	var link models.Link

	linkList, err := link.GetAll(userId.(primitive.ObjectID))
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
		ctx.Redirect(http.StatusMovedPermanently, l.Url)
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
	}
}

func deleteUrl(ctx *gin.Context) {

	id := ctx.Param("shortId")
	if strings.HasPrefix(id, "U") {
		l, err := models.GetSingleAndIncreaseClicks(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
		}
		err = l.Delete()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "failed to delete the link!"})
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted successfully!"})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
	}
}
