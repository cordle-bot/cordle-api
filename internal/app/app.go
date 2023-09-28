package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cordle-bot/cordle-api/internal/database"
	"github.com/cordle-bot/cordle-api/internal/middleware"
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

	go shutdown()

	err := godotenv.Load()
	util.ErrOut(err)

	log.Println("Connecting to DB")
	s = database.New(database.MakeSQLiteDb())

	log.Println("Starting Router")
	r = router.New()
	r.Use(middleware.MakeAuth())
	r.Use(cors.Default())
	r.RegisterRoutes(makeRoutes())
	r.NoRoute(reverseProxy())
	log.Println("Running Router")
	r.Run()
}

func shutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("")
	log.Println("Shutting down app")

	log.Println("Closing db conn")
	err := s.Close()
	util.ErrLog(err)

	os.Exit(0)
}
