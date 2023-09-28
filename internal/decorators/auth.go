package decorators

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminOnly() DecorateFunc {
	return func(fn gin.HandlerFunc) gin.HandlerFunc {
		return func(c *gin.Context) {
			h := c.GetHeader("Authorization")
			if h != "Admin" {
				c.Status(http.StatusUnauthorized)
				return
			}

			fn(c)
		}
	}
}
