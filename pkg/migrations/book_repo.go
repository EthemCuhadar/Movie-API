package database

import (
	"github.com/EthemCuhadar/Movie-API/pkg/models"
	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (m *MovieRepository) Migrations() {
	m.db.AutoMigrate(&models.Movie{})
}
