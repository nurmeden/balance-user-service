package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	TransactionId uuid.UUID `json:"transaction_id" db:"transaction_id"`
	SenderId      uuid.UUID `json:"sender_id" db:"sender_id"`
	ReceiverId    uuid.UUID `json:"receiver_id" db:"receiver_id"`
	Amount        int64     `json:"amount" db:"amount"`
	Type          string    `json:"type" db:"type"`
	Description   string    `json:"description" db:"description"`
	Timestamp     time.Time `json:"timestamp" db:"timestamp"`
	Status        string    `json:"status" db:"status"`
}
