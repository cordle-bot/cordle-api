package middleware

import "github.com/gin-gonic/gin"

func MakeAuth() gin.HandlerFunc {
	auth := gin.BasicAuth(
		gin.Accounts{
			"username": "password",
		},
	)

	return auth
}
