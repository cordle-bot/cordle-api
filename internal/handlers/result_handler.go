package handlers

import (
	"net/http"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/gin-gonic/gin"
)

// /result/win/:guild/:winner/:loser
func ResultPostWin(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		w := c.Param("winner")
		l := c.Param("loser")

		if !checkPlayers(s, w, l, guild) {
			c.Status(http.StatusBadRequest)
			return
		}

		err := calculateWin(s, w, l, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// /result/loss/:guild/:loser/:winner
func ResultPostLoss(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		l := c.Param("loser")
		w := c.Param("winner")

		if !checkPlayers(s, w, l, guild) {
			c.Status(http.StatusBadRequest)
			return
		}

		err := calculateWin(s, w, l, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

// /result/draw/:guild/:one/:two
func ResultPostDraw(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		o := c.Param("one")
		t := c.Param("two")

		if !checkPlayers(s, o, t, guild) {
			c.Status(http.StatusBadRequest)
			return
		}

		err := calculateDraw(s, o, t, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusNoContent)
	}
}

func checkPlayers(s *database.Store, o string, t string, g string) bool {
	e := s.CheckUser(o, g)
	if !e {
		return e
	}

	e = s.CheckUser(t, g)
	if !e {
		return e
	}

	return true
}

func calculateWin(s *database.Store, w string, l string, g string) error {
	return nil
}

func calculateDraw(s *database.Store, o string, t string, g string) error {
	return nil
}