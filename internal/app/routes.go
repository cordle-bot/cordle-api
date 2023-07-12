package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/CordleBot/cordle-api/internal/router"

	"github.com/CordleBot/cordle-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// The makeRoutes method returns a pointer to a Routes struct,
// which stores a slice of Route.
func makeRoutes() *router.Routes {
	return &router.Routes{
		RouteInfo: []router.Route{
			// Ping Handlers
			{
				Name:        "GetPing",
				Method:      router.Get,
				Path:        "/ping",
				Handler:     "",
				HandlerFunc: handlers.PingGet(),
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
