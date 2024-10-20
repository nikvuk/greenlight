package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`                // unique integer ID for the movie
	CreatedAt time.Time `json:"-"`                 // timestamp when the movie is added to the DB
	Title     string    `json:"title"`             // movie title
	Year      int32     `json:"year,omitempty"`    // movie release year
	Runtime   int32     `json:"runtime,omitempty"` // movie runtime (minutes)
	Genres    []string  `json:"genres,omitempty"`  // genres for the movie (romance, comedy, etc)
	Version   int32     `json:"version"`           // version number, incremented every time movie information updated
}
