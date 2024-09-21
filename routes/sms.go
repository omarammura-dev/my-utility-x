package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myutilityx.com/models"
)

func saveSms(ctx *gin.Context) {
	var sms models.SMS
	err := ctx.ShouldBindJSON(&sms)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to bind!" + err.Error()})
		return
	}

	err = sms.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "failed to save SMS! " + err.Error()})
		return
	}
}
func getSms(ctx *gin.Context) {
	var sms models.SMS
	err := sms.Get()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error while finding the sms"})
	}
	ctx.JSON(http.StatusOK, sms)
	sms.Delete()
}
