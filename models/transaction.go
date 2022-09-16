package models

import "time"

type Transaction struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    string    `json:"status"`
	UserID    int       `json:"user_id" gorm:"type:int"`
	User      User      `json:"user"`
}

type TransactionRelation struct {
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    bool      `json:"status"`
}

func (TransactionRelation) TableName() string {
	return "transactions"
}
