package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(_db *gorm.DB) *Repository {

	_repository := Repository{db: _db}

	return &_repository
}