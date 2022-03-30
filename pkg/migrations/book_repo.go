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

func (m *MovieRepository) InsertData(list models.MovieList) {
	for _, movie := range list {
		m.db.Where(models.Movie{ID: movie.ID}).
			Attrs(models.Movie{ID: movie.ID, Title: movie.Title, Year: movie.Year, Runtime: movie.Runtime,
				Genres: movie.Genres, Director: movie.Director, Actors: movie.Actors, Plot: movie.Plot, PosterURL: movie.PosterURL}).
			FirstOrCreate(&movie)
	}
}
