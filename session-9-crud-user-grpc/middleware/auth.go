package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ibrahimker/golang-praisindo-advanced/session-9-crud-user-grpc/config"
	"net/http"
)

// AuthMiddleware adalah middleware untuk autentikasi
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifikasi token (misalnya, cocokkan dengan token yang diharapkan)
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization basic token required"})
			c.Abort()
			return
		}

		isValid := (username == config.AuthBasicUsername) && (password == config.AuthBasicPassword)
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika token valid
		c.Next()
	}
}
