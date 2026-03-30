package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/joshhartwig/movies/data"
	"github.com/joshhartwig/movies/handlers"
	"github.com/joshhartwig/movies/logger"
	_ "github.com/lib/pq"
)

func initLogger() *logger.Logger {
	logger, err := logger.NewLogger("logs/movies.log")
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Close()
	return logger
}

func main() {
	logger := initLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatal("No env file available")
	}

	connString := os.Getenv("DB_URL")
	if connString == "" {
		log.Fatal("Error finding db connection string")
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Failed to connected to database with error: %v", err)
	}

	defer db.Close()

	// init data repository
	movieRepo, err := data.NewMovieRepository(db, logger)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	const addr = ":8080"

	movieHandler := handlers.NewMovieRepository(movieRepo, logger)

	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)
	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)
	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)
	http.HandleFunc("/api/movies/", movieHandler.GetMovie)
	http.HandleFunc("/api/genres", movieHandler.GetGenres)
	http.HandleFunc("/api/account/register", movieHandler.GetGenres)
	http.HandleFunc("/api/account/authenticate", movieHandler.GetGenres)

	http.Handle("/", http.FileServer(http.Dir("public")))

	logger.Info(fmt.Sprintf("starting server on port %s", addr))
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("error with server %v", err)
		logger.Error("error", err)
	}

}
