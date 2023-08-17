package database

import (
	"github.com/cordle-bot/cordle-api/internal/models"
	"gorm.io/gorm/clause"
)

func (s *Store) GetLeaderboard(t string) ([]models.UserModel, error) {
	m := make([]models.UserModel, 0)
	r := s.db.Table(t).Order(clause.OrderByColumn{Column: clause.Column{Name: "elo"}, Desc: true}).Limit(10).Find(&m)

	return m, r.Error
}
