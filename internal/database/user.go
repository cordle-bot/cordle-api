package database

import "github.com/cordle-bot/cordle-api/internal/models"

func (s *Store) ListUsers(t string) ([]models.UserModel, error) {
	m := make([]models.UserModel, 0)
	r := s.db.Table(t).Find(&m)
	return m, r.Error
}

func (s *Store) GetUser(i string, t string) (models.UserModel, error) {
	var m models.UserModel
	r := s.db.Table(t).Find(&m, i)
	return m, r.Error
}

func (s *Store) AddUser(m models.UserModel, t string) error {
	r := s.db.Table(t).Create(&m)
	return r.Error
}

func (s *Store) UpdateUser(m models.UserModel, t string) error {
	r := s.db.Table(t).Save(&m)
	return r.Error
}

func (s *Store) DeleteUser(i string, t string) error {
	p := models.UserModel{Id: i}
	r := s.db.Table(t).Delete(&p)
	return r.Error
}

func (s *Store) CheckUser(i string, t string) bool {
	m := models.UserModel{Id: i}
	r := s.db.Table(t).Model(&m).First(&m)
	return r.Error == nil
}
