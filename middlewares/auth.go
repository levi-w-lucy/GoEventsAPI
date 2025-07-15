package middlewares

import (
	"net/http"

	"example.com/EventsAPI/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	parsedToken, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.GetUserIDClaim(parsedToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid user id."})
		return
	}

	context.Set("UserID", userId)

	context.Next()
}
