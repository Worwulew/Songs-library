package model

import "time"

// Song is struct for song
// @Description Song is a struct that represents a song
type Song struct {
	SongID      uint      `json:"song_id" db:"song_id" validate:"omitempty" swaggerignore:"true"`
	Group       string    `json:"group" db:"group_name" validate:"required" example:"AC/DC"`
	SongTitle   string    `json:"songTitle" db:"song_title" validate:"required" example:"Long way to Hell"`
	ReleaseDate string    `json:"releaseDate" db:"release_date" validate:"omitempty"`
	Text        string    `json:"text" db:"text" validate:"omitempty"`
	Link        string    `json:"link" db:"link" validate:"omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at" format:"date-time" swaggerignore:"true"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at" format:"date-time" swaggerignore:"true"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate" db:"release_date" validate:"omitempty"`
	Text        string `json:"text" db:"text" validate:"omitempty"`
	Link        string `json:"link" db:"link" validate:"omitempty"`
}

// SongsList is all Song response
// @Description SongsList is a struct list of songs for searching response
type SongsList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Songs      []*Song `json:"songs"`
}

// VersesList is struct for getting verses
// @Description VersesList is a struct list of verses of song text
type VersesList struct {
	TotalCount int      `json:"totalCount"`
	TotalPages int      `json:"totalPages"`
	Page       int      `json:"page"`
	Size       int      `json:"size"`
	HasMore    bool     `json:"hasMore"`
	Verses     []string `json:"verses"`
}
