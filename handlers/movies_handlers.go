package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/joshhartwig/movies/data"
	"github.com/joshhartwig/movies/logger"
	"github.com/joshhartwig/movies/models"
)

type MovieHandler struct {
	logger  *logger.Logger
	storage data.MovieStorage
}

func NewMovieRepository(storage data.MovieStorage, logger *logger.Logger) *MovieHandler {
	return &MovieHandler{
		storage: storage,
		logger:  logger,
	}
}

func (m *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.storage.GetTopMovies()
	if err != nil {
		m.logger.Error("Error fetching movies from repository", err)
		http.Error(w, "error fetching movies from repository", http.StatusInternalServerError)
		return
	}

	m.writeJSONResponse(w, &movies)
}

func (m *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.storage.GetRandomMovies()
	if err != nil {
		m.logger.Error("Error fetching movies from repository", err)
		http.Error(w, "error fetching movies from repository", http.StatusInternalServerError)
		return
	}

	m.writeJSONResponse(w, &movies)
}

// Utility functions
func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		m.logger.Error("Error fetching the movie id", errors.New("invalid id"))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	parsedID, err := strconv.Atoi(id)
	if err != nil {
		m.logger.Error("Error fetching the movie id", errors.New("invalid id"))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	movie, err := m.storage.GetMovieByID(parsedID)
	if err != nil {
		m.logger.Error("Error fetching the movie with that id", errors.New("invalid id"))
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	m.writeJSONResponse(w, movie)
}

func (m *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := m.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = m.storage.SearchMoviesByName(query, order, genre)
	}
	if m.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if m.writeJSONResponse(w, movies) == nil {
		m.logger.Info("Successfully served movies")
	}
}

func (h *MovieHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == data.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.storage.GetAllGenres()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	if h.writeJSONResponse(w, genres) == nil {
		h.logger.Info("Successfully served genres")
	}
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/api/movies/"):]
	id, ok := m.parseID(w, idStr)
	if !ok {
		return
	}

	movie, err := m.storage.GetMovieByID(id)
	if m.handleStorageError(w, err, "Failed to get movie by ID") {
		return
	}
	if m.writeJSONResponse(w, movie) == nil {
		m.logger.Info("Successfully served movie with ID: " + idStr)
	}
}
