package authdto

type LoginResponse struct {
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	ID      int    `gorm:"type: int" json:"id"`
	Name    string `gorm:"type: varchar(255)" json:"name"`
	Email   string `gorm:"type: varchar(255)" json:"email"`
	IsAdmin bool   `gorm:"type: boolean" json:"isAdmin"`
}
