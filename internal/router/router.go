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
func New() *Router {
	return &Router{
		Engine: gin.Default(),
	}
}

// Wrapper for the gin engine NoRoute method.
func (r *Router) NoRoute(f gin.HandlerFunc) {
	r.Engine.NoRoute(f)
}

// Wrapper for the gin engine Use method.
func (r *Router) Use(f gin.HandlerFunc) {
	r.Engine.Use(f)
}

// Wrapper for the gin engine Run method.
func (r *Router) Run() {
	r.Engine.Run()
}

// Wrapper for the gin engine GET method.
func (r *Router) Get(route Route) {
	r.Engine.GET(route.Path+route.Handler, route.HandlerFunc)
}

// Wrapper for the gin engine POST method.
func (r *Router) Post(route Route) {
	r.Engine.POST(route.Path+route.Handler, route.HandlerFunc)
}

// Wrapper for the gin engine PUT method.
func (r *Router) Put(route Route) {
	r.Engine.PUT(route.Path+route.Handler, route.HandlerFunc)
}

// Wrapper for the gin engine PATCH method.
func (r *Router) Patch(route Route) {
	r.Engine.PATCH(route.Path+route.Handler, route.HandlerFunc)
}

// Wrapper for the gin engine DELETE method.
func (r *Router) Delete(route Route) {
	r.Engine.DELETE(route.Path+route.Handler, route.HandlerFunc)
}

// Registers a route.
//
// If the route method is undefined it will cause a panic.
func (r *Router) RegisterRoute(route Route) {
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

// Registers all routes passed depending on there method type.
//
// If the route method is undefined it will cause a panic.
func (r *Router) RegisterRoutes(routes Routes) {
	for _, route := range routes.RouteInfo {
		r.RegisterRoute(route)
	}
}
