package models

import "time"

type User struct {
	ID           int             `json:"id" gorm:"primary_key:auto_increment"`
	Email        string          `json:"email" gorm:"type: varchar(255)"`
	Password     string          `json:"-" gorm:"type: varchar(255)"`
	IsAdmin      bool            `json:"isAdmin"`
	Profile      ProfileResponse `json:"profile" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transactions []Transaction   `json:"transactions"`
	Films        []Film          `json:"film"`
	CreatedAt    time.Time       `json:"-"`
	UpdatedAt    time.Time       `json:"-"`
}

type UsersProfileResponse struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}

type UserTransaction struct {
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    bool      `json:"status"`
}

type UserFilm struct {
	Title         string `json:"title"`
	ThumbnailFilm string `json:"thumbnailfilm"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	Year          int    `json:"year"`
}

func (UserFilm) TableName() string {
	return "users"
}

func (UserTransaction) TableName() string {
	return "users"
}

func (UsersProfileResponse) TableName() string {
	return "users"
}
