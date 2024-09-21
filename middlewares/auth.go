package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myutilityx.com/utils"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized (empty)"})
		return
	}

	userId, role, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Set("userId", userId)
	ctx.Set("role", role)
	ctx.Next()
}
