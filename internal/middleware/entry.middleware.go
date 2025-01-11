package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sketch873/go-departments-project/internal/api"
)
func EntryMiddleware(provider *api.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("provider", provider)
		c.Next()
	}
}