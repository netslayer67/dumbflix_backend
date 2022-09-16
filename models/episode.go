package models

import "time"

type Episode struct {
	ID            int       `json:"id" gorm:"primary_key:auto_increment"`
	Title         string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string    `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	LinkFilm      string    `json:"linkfilm" gorm:"type:text" form:"linkfilm"`
	FilmID        int       `json:"-"`
	Film          Film      `json:"film"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type EpisodeResponse struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	LinkFilm      string `json:"linkfilm"`
	FilmID        int    `json:"-"`
	Film          Film   `json:"film"`
}

func (EpisodeResponse) TableName() string {
	return "episodes"
}
