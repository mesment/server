package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mesment/server/handler"
	"github.com/mesment/server/pkg/errno"
	"github.com/mesment/server/pkg/token"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
