package userstorage

import (
	"gorm.io/gorm"
)

// Encapsulation
type sqlStore struct {
	db *gorm.DB
}

func NewSqlStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db: db}
}
