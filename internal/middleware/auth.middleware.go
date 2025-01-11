package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sketch873/go-departments-project/internal/session"
	"github.com/sketch873/go-departments-project/pkg/mysql/user"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := session.GetSession(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		user, err := user.GetUserByUuid(uuid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		if user == nil {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}