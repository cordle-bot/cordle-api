package router

import (
	"github.com/cordle-bot/cordle-api/internal/decorators"
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
	Name          string                  // Route name
	Method        Method                  // Route method
	Path          string                  // Route path
	Params        string                  // Route Handler
	HandlerFunc   gin.HandlerFunc         // Route Handler Function
	DecoratorFunc decorators.DecorateFunc // Route Decorator Function
}

// Routes stores a slice of the Route struct.
//   - RouteInfo : slice of Route structs.
type Routes struct {
	RouteInfo []Route // Routes-
}
