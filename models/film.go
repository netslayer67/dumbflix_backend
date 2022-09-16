package models

import "time"

type Film struct {
	ID            int       `json:"id" gorm:"primary_key:auto_increment"`
	Title         string    `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string    `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Image         string    `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description   string    `json:"description" gorm:"type:text" form:"desc"`
	Year          int       `json:"year" gorm:"type: int"`
	CategoryID    int       `json:"-" gorm:"type:int"`
	Category      Category  `json:"category"`
	UserID        int       `json:"-" gorm:"type:int"`
	User          User      `json:"user"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}

type UserRelation struct {
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	Year          int    `json:"year"`
}

type FilmRelation struct {
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnail"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	Year          int    `json:"year"`
	CategoryID    int    `json:"category_id"`
}

func (FilmRelation) TableName() string {
	return "films"
}
