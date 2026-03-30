package data

import "github.com/joshhartwig/movies/models"

type MovieStorage interface {
	GetTopMovies() ([]models.Movie, error)
	GetMovieByID(id int) (models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	SearchMoviesByName(name string, order string, genre *int) ([]models.Movie, error)
	GetAllGenres() ([]models.Genre, error)
}
