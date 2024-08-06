package data

import "time"

type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genres      []string  `json:"genres,omitempty"`
	CreatedAt   time.Time `json:"-"`
}
