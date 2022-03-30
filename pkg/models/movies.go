package models

// "encoding/json"

// Movie model with relative fields created. Necessary JSON
// parameters can be seen below.
type Movie struct {
	ID        int      `json:"id" gorm:"primary_key;unique;type:INT;not null"`
	Title     string   `json:"title" gorm:"size:100;not null"`
	Year      string   `json:"year"`
	Runtime   string   `json:"runtime"`
	Genres    []string `json:"genres"`
	Director  string   `json:"director"`
	Actors    string   `json:"actors"`
	Plot      string   `json:"plot"`
	PosterURL string   `json:"posterUrl"`
}

// MovieList created to store movie instances.
type MovieList []Movie
