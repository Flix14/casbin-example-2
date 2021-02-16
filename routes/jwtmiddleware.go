package routes

import (
	"net/http"

	m "github.com/Flix14/casbin-example-2/models"
	"github.com/gin-gonic/gin"
)

func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")

		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claim, err := m.ValidateTokenJWT(token, c.FullPath(), c.Request.Method)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("USER", claim.User)
		c.Next()
	}
}
