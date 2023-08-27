package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/cordle-bot/cordle-api/internal/middleware"
	"github.com/cordle-bot/cordle-api/internal/models"
	"github.com/cordle-bot/cordle-api/internal/router"
	"github.com/cordle-bot/cordle-api/pkg/util"
	"github.com/gin-contrib/cors"

	"github.com/joho/godotenv"
)

// Stores the router
var r *router.Router

// Stores the DB
var s *database.Store

// Runs the app
//   - Creates the router.
//   - Registers routes.
//   - Runs the engine.
//   - Waits for interrupt before shutting down app.
func Run() {
	log.Println("Starting app")

	err := godotenv.Load()
	util.ErrOut(err)

	s = database.New(database.MakeSQLiteDb())

	r = router.New()
	r.Use(middleware.MakeAuth())
	r.RegisterRoutes(makeRoutes())
	r.Use(cors.Default())
	r.NoRoute(reverseProxy())
	r.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down app")
}
