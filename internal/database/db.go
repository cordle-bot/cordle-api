package database

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/cordle-bot/cordle-api/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Storers a *gorm.DB and its mutex.
type Storer struct {
	db   *gorm.DB
	dbMu sync.Mutex
}

// Make a new Storer.
//
// *gorm.DB is made from .env - follow .env.example pattern.
func New() *Storer {
	return &Storer{
		db: makeDb(),
	}
}

// Ping the storer.
//
// If not connected an error will be returned.
func (s *Storer) Ping() error {
	s.dbMu.Lock()
	defer s.dbMu.Unlock()

	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}

// Creates a pointer to a gorm db.
//
// This uses environmental variables for the dsn.
//
// A connection is then opened, checked for errors and returned.
//
// Keys for environmental variables:
//   - DB_ADDR : stores the address (IP)
//   - DB_PORT : stores the port
//   - DB_USER : stores the username
//   - DB_PASS : stores the password
//   - DB_NAME : stores the database name
func makeDb() *gorm.DB {
	dsn := fmt.Sprintf(`
		host=%s 
		user=%s 
		password=%s 
		dbname=%s 
		port=%s 
		sslmode=disable 
		TimeZone=Europe/London`,
		os.Getenv("DB_ADDR"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	util.ErrOut(err)

	return db
}
