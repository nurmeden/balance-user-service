package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	OrderID   uuid.UUID `json:"order_id" db:"order_id"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	ServiceID uuid.UUID `json:"service_id" db:"service_id"`
	Amount    float64   `json:"amount" db:"amount"`
	OrderDate time.Time `json:"order_date" db:"order_date"`
	Status    string    `json:"status" db:"status"`
}
