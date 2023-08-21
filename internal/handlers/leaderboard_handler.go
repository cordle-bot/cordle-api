package handlers

import (
	"net/http"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/cordle-bot/cordle-api/internal/models"
	"github.com/gin-gonic/gin"
)

// /leaderboard/:guild
func LeaderboardGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")

		l, err := s.GetLeaderboard(guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r := make([]models.UserPost, len(l))
		for i, user := range l {
			r[i] = user.ToPost()
		}

		c.JSON(http.StatusOK, r)
	}
}
