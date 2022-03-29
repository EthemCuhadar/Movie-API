package models

// "encoding/json"

type Movie struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Year      string   `json:"year"`
	Runtime   string   `json:"runtime"`
	Genres    []string `json:"genres"`
	Director  string   `json:"director"`
	Actors    string   `json:"actors"`
	Plot      string   `json:"plot"`
	PosterURL string   `json:"posterUrl"`
}

type MovieList []Movie
