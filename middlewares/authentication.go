package middlewares

import (
	"final-project/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			result := gin.H{
				"status":  false,
				"message": "Unauthenticated",
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, result)
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
