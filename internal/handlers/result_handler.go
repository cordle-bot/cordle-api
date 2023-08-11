package handlers

import "github.com/gin-gonic/gin"

func ResultGet() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ResultPost() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ResultPut() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ResultPatch() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func ResultDelete() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /result/win/:winner/:loser
func ResultPostWin() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /result/loss/:loser/:winner
func ResultPostLoss() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// /result/draw/:player/:player
func ResultPostDraw() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func calculateResult(w string, l string) error {
	return nil
}
