package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Creates a gin HandlerFunc.
//
// Example usage:
//   - /ping : returns {"ping":"pong"}.
func PingGet() gin.HandlerFunc {
	resp := "pong"

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": resp,
		})
	}
}
