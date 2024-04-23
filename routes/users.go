package routes

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"myutilityx.com/mailS"
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

	token, err := utils.GenerateToken(user.Email, user.Username, user.ID, time.Hour*2)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "OOps, Something went wrong!"})
	}
	_, err = mailS.SendSimpleMessage(os.Getenv("API_URL")+"user/verify?token="+token, user.Email,user.Username,"d-958c75cdb588424fb80e49688fb2c3da")
	if err != nil {
		log.Fatalf("could not send verification email: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not send verification email"})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func login(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while binding the user data"})
		return
	}


	err = user.ValidateCredintials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "" + err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Username, user.ID, time.Hour*2)

	if err != nil {
		log.Fatalf("could not generate the token: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate the token"})
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user, "token": token})
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

	err = user.Update(bson.M{"isverified": true})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "oops!" + err.Error()})

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "user is verified."})
}

func resetPassword(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if user.Email == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email is empty!"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong (643)! " + err.Error()})
		return
	}

	err = user.FindByEmail()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There seems to be a discrepancy with the information provided. To ensure account security, we can't assist with password resets for unrecognized information."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.Username, user.ID, time.Minute*15)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong (732)!"})
		return
	}

	_, err = mailS.SendSimpleMessage(os.Getenv("UI_URL")+"auth/reset-password/confirm/"+token, user.Email,user.Username,"d-325e3a95b2fb497d9c293519596f6a45")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email could not send!"})
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"Reset password email sent successfully! please check your inbox."})
}   

func resetPasswordVerify(ctx *gin.Context) {
	token := ctx.Query("token")
	var user models.User
	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong (732)!"})
    return
	}

	var passwords struct {
		OldPassword string `binding:"required"`
		NewPassword string `binding:"required"`
	}

	err := ctx.ShouldBindJSON(&passwords)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not parse!"})
		return
	}


	userid, err := utils.VerifyToken(token)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "OOPs!" + err.Error()})
		return
	}

	user.ID = userid

	err = user.VerifyAndUpdatePassword(passwords.OldPassword)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Please double check your password!"})	
		return
	}

	hashedPassword,err := utils.HashPassword(passwords.NewPassword)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "error!" + err.Error()})	
		return
	}

	err = user.Update(bson.M{"password":hashedPassword})

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "error!" + err.Error()})	
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password" + userid.String()})
}
