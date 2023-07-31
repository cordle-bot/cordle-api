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
				Handler:     "",
				HandlerFunc: handlers.PingGet(),
			},
			// Guild Handlers
			{
				Name:        "Guild Get",
				Method:      router.Get,
				Path:        "/guild",
				Handler:     "/:guild",
				HandlerFunc: handlers.GuildGet(),
			},
			// User Handlers
			{
				Name:        "User Get",
				Method:      router.Get,
				Path:        "/user",
				Handler:     "/:user",
				HandlerFunc: handlers.UserGet(),
			},
			{
				Name:        "User Post",
				Method:      router.Post,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPost(),
			},
			{
				Name:        "User Put",
				Method:      router.Put,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPut(),
			},
			{
				Name:        "User Patch",
				Method:      router.Patch,
				Path:        "/user",
				Handler:     "",
				HandlerFunc: handlers.UserPatch(),
			},
			{
				Name:        "User Delete",
				Method:      router.Delete,
				Path:        "/user",
				Handler:     "/:user",
				HandlerFunc: handlers.UserDelete(),
			},
			// Leaderboard Handlers
			{
				Name:        "Leaderboard Get",
				Method:      router.Get,
				Path:        "/leaderboard",
				Handler:     "",
				HandlerFunc: handlers.LeaderboardGet(),
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
