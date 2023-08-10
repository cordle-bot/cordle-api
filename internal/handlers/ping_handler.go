package handlers

import (
	"net/http"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/gin-gonic/gin"
)

// Creates a gin HandlerFunc.
//
// Example usage:
//   - /ping : returns {"ping":"pong"}.
func PingGet(s *database.Storer) gin.HandlerFunc {
	resp := "pong"
	err := s.Ping()
	if err != nil {
		resp = "err"
	}

	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": resp,
		})
	}
}
