package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"myutilityx.com/errors"
	"myutilityx.com/models"
)

func addLink(ctx *gin.Context) {

	link, err := models.InitLink()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errors.ErrSomethingWentWrong)
		return
	}
	err = ctx.ShouldBindJSON(&link)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse the link object!"})
		return
	}

	userId, exist := ctx.Get("userId")
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

	userId, exist := ctx.Get("userId")

	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Oops! something went wrong!"})
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
	id := ctx.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID format"})
		return
	}

	l, err := models.FindById(objectID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Not found!"})
		return
	}

	err = l.Delete()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the link!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted successfully!"})
}
