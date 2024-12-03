package middleware

import (
	"idler/app/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = http.StatusOK

		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusUnauthorized
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = http.StatusInternalServerError
			}
			c.Set("claims", claims)
		}

		if code != http.StatusOK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
