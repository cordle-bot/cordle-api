package router

import (
	"github.com/gin-gonic/gin"
)

// Route struct stores information for create a route in the gin engine.
//   - Name : stores the name of the route.
//   - Method : what type is it? Post, Patch, Delete, etc.?
//   - Path : path of the route.
//   - Handler : handlers of the route.
//   - HandlerFunc : gin.Handlerfunc,
//     stores the method that occurs when this route is queried.
type Route struct {
	Name        string          // Route name
	Method      Method          // Route method
	Path        string          // Route path
	Handler     string          // Route Handler
	HandlerFunc gin.HandlerFunc // Route Handler Function
}

// Routes stores a slice of the Route struct.
//   - RouteInfo : slice of Route structs.
type Routes struct {
	RouteInfo []Route // Routes-
}
