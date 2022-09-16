package categoriesdto

type CategoryResponse struct {
	ID    int            `json:"id" gorm:"primary_key:auto_increment"`
	Name  string         `json:"name"`
	Films []CategoryFilm `json:"films"`
}

type CategoryFilm struct {
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Description   string `json:"desc"`
	Year          int    `json:"year"`
}
