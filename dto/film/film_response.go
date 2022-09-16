package filmsdto

type FilmResponse struct {
	ID            int                  `json:"id" gorm:"primary_key:auto_increment"`
	Title         string               `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string               `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Description   string               `json:"desc" gorm:"type:text" form:"desc"`
	Year          int                  `json:"year" form:"year" gorm:"type: int"`
	Category      CategoryFilmResponse `json:"category"`
}

type CategoryFilmResponse struct {
	Name string `json:"name"`
}
