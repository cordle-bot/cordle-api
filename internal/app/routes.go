package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/cordle-bot/cordle-api/internal/router"

	"github.com/cordle-bot/cordle-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// The makeRoutes method returns a pointer to a Routes struct,
// which stores a slice of Route.
func makeRoutes() router.Routes {
	return router.Routes{
		RouteInfo: []router.Route{
			// Ping Handlers
			{
				Name:        "Ping Get",
				Method:      router.Get,
				Path:        "/ping",
				Params:      "",
				HandlerFunc: handlers.PingGet(s),
			},
			// Guild Handlers
			{
				Name:        "Guild Get",
				Method:      router.Get,
				Path:        "/guild",
				Params:      "/:guild",
				HandlerFunc: handlers.GuildGet(),
			},
			// User Handlers
			{
				Name:        "User List",
				Method:      router.Get,
				Path:        "/user",
				Params:      "/:guild",
				HandlerFunc: handlers.UserList(s),
			},
			{
				Name:        "User Get",
				Method:      router.Get,
				Path:        "/user",
				Params:      "/:guild/:user",
				HandlerFunc: handlers.UserGet(s),
			},
			{
				Name:        "User Post",
				Method:      router.Post,
				Path:        "/user",
				Params:      "/:guild",
				HandlerFunc: handlers.UserPost(s),
			},
			{
				Name:        "User Put",
				Method:      router.Put,
				Path:        "/user",
				Params:      "/:guild",
				HandlerFunc: handlers.UserPut(s),
			},
			{
				Name:        "User Patch",
				Method:      router.Patch,
				Path:        "/user",
				Params:      "/:guild",
				HandlerFunc: handlers.UserPatch(s),
			},
			{
				Name:        "User Delete",
				Method:      router.Delete,
				Path:        "/user",
				Params:      "/:guild/:user",
				HandlerFunc: handlers.UserDelete(s),
			},
			// Leaderboard Handlers
			{
				Name:        "Leaderboard Get",
				Method:      router.Get,
				Path:        "/leaderboard",
				Params:      "/:guild",
				HandlerFunc: handlers.LeaderboardGet(s),
			},
			// Result Handlers
			{
				Name:        "Result Win Post",
				Method:      router.Post,
				Path:        "/result/win",
				Params:      "/:winner/:loser",
				HandlerFunc: handlers.ResultPostWin(s),
			},
			{
				Name:        "Result Loss Post",
				Method:      router.Post,
				Path:        "/result/loss",
				Params:      "/:loser/:winner",
				HandlerFunc: handlers.ResultPostLoss(s),
			},
			{
				Name:        "Result Draw Post",
				Method:      router.Post,
				Path:        "/result/draw",
				Params:      "/:one/:two",
				HandlerFunc: handlers.ResultPostDraw(s),
			},
		},
	}
}

// If gin cannot find it use this proxy.
func reverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		remote, _ := url.Parse("http://localhost:3000")
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL = c.Request.URL
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
