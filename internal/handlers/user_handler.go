package handlers

import (
	"net/http"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/cordle-bot/cordle-api/internal/models"
	"github.com/gin-gonic/gin"
)

// /user/:guild
func UserList(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")

		l, err := s.ListUsers(guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		r := make([]models.UserPost, 0)
		for _, user := range l {
			r = append(r, user.ToPost())
		}

		c.JSON(http.StatusOK, r)
	}
}

// /user/:guild/:user
func UserGet(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		user := c.Param("user")

		e := s.CheckUser(user, guild)
		if !e {
			c.Status(http.StatusNotFound)
			return
		}

		g, err := s.GetUser(user, guild)
		r := g.ToPost()

		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, r)
	}
}

// /user/:guild
func UserPost(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		b := models.UserPost{}
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckUser(m.Id, guild)
		if e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.AddUser(m, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// /user/:guild
func UserPut(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		b := models.UserPost{}
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckUser(m.Id, guild)
		if !e {
			err := s.AddUser(m, guild)
			if err != nil {
				c.Status(http.StatusBadRequest)
				return
			}

			c.Status(http.StatusOK)
			return
		}

		err := s.UpdateUser(m, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// /user/:guild
func UserPatch(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		b := models.UserPost{}
		c.ShouldBindJSON(&b)
		m := b.ToModel()

		e := s.CheckUser(m.Id, guild)
		if !e {
			c.Status(http.StatusBadRequest)
			return
		}

		err := s.UpdateUser(m, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}

// /user/:guild/:user
func UserDelete(s *database.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		guild := c.Param("guild")
		user := c.Param("user")

		e := s.CheckUser(user, guild)
		if !e {
			c.Status(http.StatusNotFound)
		}

		err := s.DeleteUser(user, guild)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return
		}

		c.Status(http.StatusOK)
	}
}
