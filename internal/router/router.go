package router

import (
	"github.com/cordle-bot/cordle-api/pkg/util"

	"github.com/gin-gonic/gin"
)

// Router stores pointers of both a gin engine and a store.
//   - Engine : gin engine that is responsible for routing.
//   - Store : using gorm retrieves information for the engine.
type Router struct {
	Engine *gin.Engine
}

// Returns a pointer to a default router.
//   - Engine : default gin router.
//   - Store : default store using env variables for db connection.
func MakeRouter() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

func (r *Router) NoRoute(f gin.HandlerFunc) {
	r.Engine.NoRoute(f)
}

func (r *Router) Use(f gin.HandlerFunc) {
	r.Engine.Use(f)
}

// Wrapper for the gin engine Run method.
func (r *Router) Run() {
	r.Engine.Run()
}

// Registers all routes passed depending on there method type.
//
// If the route method is undefined it will cause a panic.
func (r *Router) RegisterRoutes(rl Routes) {
	for _, route := range rl.RouteInfo {
		switch route.Method {
		case Undefined:
			util.ErrOut(util.ErrUndefinedRouteMethod)
		case Get:
			r.Engine.GET(route.Path+route.Handler, route.HandlerFunc)
		case Post:
			r.Engine.POST(route.Path+route.Handler, route.HandlerFunc)
		case Put:
			r.Engine.PUT(route.Path+route.Handler, route.HandlerFunc)
		case Patch:
			r.Engine.PATCH(route.Path+route.Handler, route.HandlerFunc)
		case Delete:
			r.Engine.DELETE(route.Path+route.Handler, route.HandlerFunc)
		}
	}
}
