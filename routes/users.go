package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"myutilityx.com/mailS"
	"myutilityx.com/models"
	"myutilityx.com/utils"
	"net/http"
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



	token, err := utils.GenerateToken(user.Email, user.Username, user.ID)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "OOps, Something went wrong!"})
	}
	_, err = mailS.SendSimpleMessage("http://localhost:8080/register/verify?token="+token, "omarammoralm10@gmail.com", "omarammura")
	if err != nil {
		log.Fatalf("could not send verification email: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not send verification email"})
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
	ctx.JSON(http.StatusOK, gin.H{"message": "logged in", "token": token})
}

func verifyEmail(ctx *gin.Context) {
	token, ok := ctx.GetQuery("token")

	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Incomplete request!"})
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect or expired token!"})
	}

	var user models.User

	user.ID = userId

	err = user.Update()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "oops!" + err.Error()})

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user is verified."})
}
