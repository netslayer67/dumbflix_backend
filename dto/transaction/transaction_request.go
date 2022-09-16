package transactionsdto

import "time"

type CreateTransactionRequest struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    string    `json:"category"`
}

type UpdateTransactionRequest struct {
	ID        string    `json:"id"`
	StartDate time.Time `json:"statusdate"`
	DueDate   time.Time `json:"duedate"`
	Attache   string    `json:"attache"`
	Status    bool      `json:"category"`
}
