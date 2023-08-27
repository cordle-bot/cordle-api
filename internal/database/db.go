package database

import (
	"context"
	"fmt"
	"os"

	"github.com/cordle-bot/cordle-api/internal/models"
	"github.com/cordle-bot/cordle-api/pkg/util"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	UsersTable = "users"
)

// Storers a *gorm.DB and its mutex.
type Store struct {
	db *gorm.DB
}

type DbMaker func() *gorm.DB

// Make a new Storer.
//
// *gorm.DB is made from .env - follow .env.example pattern.
func New(m DbMaker) *Store {
	return &Store{
		db: m(),
	}
}

// Close the database
func (s *Store) Close() error {
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	return err
}

// Ping the storer.
//
// If not connected an error will be returned.
func (s *Store) Ping() error {
	ctx := context.Background()
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	err = db.PingContext(ctx)
	return err
}

func (s *Store) CreateTable(n string, m interface{}) error {
	err := s.db.Table(n).AutoMigrate(m)
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
func MakePostgresDb() DbMaker {
	return func() *gorm.DB {
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
}

// Creates a pointer to a gorm db.
//
// This uses an environmental variable for the path to the sqlite db.
//
// Keys for environmental variables:
//   - SQLITE_DB : file name of sqlite db.
func MakeSQLiteDb() DbMaker {
	return func() *gorm.DB {
		db, err := gorm.Open(sqlite.Open(os.Getenv("SQLITE_DB")), &gorm.Config{})
		util.ErrOut(err)
		return db
	}
}

// Lists all users stored in the given table.
func (s *Store) ListUsers(t string) ([]models.UserModel, error) {
	m := make([]models.UserModel, 0)
	r := s.db.Table(t).Find(&m)
	return m, r.Error
}

// Returns a user with the given id and table.
func (s *Store) GetUser(i string, t string) (models.UserModel, error) {
	var m models.UserModel
	r := s.db.Table(t).Find(&m, i)
	return m, r.Error
}

// Adds the given UserModel to the given table.
func (s *Store) AddUser(m models.UserModel, t string) error {
	r := s.db.Table(t).Create(&m)
	return r.Error
}

// Updates the given UserModel in the given table.
func (s *Store) UpdateUser(m models.UserModel, t string) error {
	r := s.db.Table(t).Save(&m)
	return r.Error
}

// Deletes the given user id in the given table.
func (s *Store) DeleteUser(i string, t string) error {
	p := models.UserModel{Id: i}
	r := s.db.Table(t).Delete(&p)
	return r.Error
}

// Checks if a given user id exists in the given table.
func (s *Store) CheckUser(i string, t string) bool {
	m := models.UserModel{Id: i}
	r := s.db.Table(t).Model(&m).First(&m)
	return r.Error == nil
}

// Gets the leaderboard, the, at max, top 10 users in the table by elo.
func (s *Store) GetLeaderboard(t string) ([]models.UserModel, error) {
	m := make([]models.UserModel, 0)
	r := s.db.Table(t).Order(clause.OrderByColumn{Column: clause.Column{Name: "elo"}, Desc: true}).Limit(10).Find(&m)

	return m, r.Error
}
