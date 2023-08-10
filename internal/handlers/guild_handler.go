package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cordle-bot/cordle-api/internal/models"
	"github.com/cordle-bot/cordle-api/pkg/util"
	"github.com/gin-gonic/gin"
)

func GuildGet() gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("guild")
		url := fmt.Sprintf("https://discord.com/api/guilds/%s?with_counts=true", i)

		client := http.Client{
			Timeout: time.Second * 2,
		}
		req, err := http.NewRequest(http.MethodGet, url, nil)
		util.LogErr(err)

		req.Header.Add("Authorization", os.Getenv("DISCORD_TOK"))
		req.Header.Set("User-Agent", "cordle-api")

		res, err := client.Do(req)
		util.LogErr(err)

		if res.Body != nil {
			defer res.Body.Close()
		}

		var g models.Guild
		err = json.NewDecoder(res.Body).Decode(&g)
		util.LogErr(err)

		c.JSON(http.StatusOK, g)
	}
}
