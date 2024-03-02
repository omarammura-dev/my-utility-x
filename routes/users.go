package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"myutilityx.com/models"
	"myutilityx.com/utils"
)

func register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Fatalf("error while binding the user data: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while binding the user data"})
	}
	err = user.Save()
	if err != nil {
		log.Fatalf("could not save the user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not register the user"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user created!"})

}

func login(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		log.Fatalf("error while binding the user data: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while binding the user data"})
	}

	err = user.ValidateCredintials()

	if err != nil {
		log.Fatalf("could not login the user: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not login the user"})
	}
	
	token, err := utils.GenerateToken(user.Email, user.Username, user.ID)

	if err != nil {
		log.Fatalf("could not generate the token: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate the token"})	
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "logged in", "token":token})
}
