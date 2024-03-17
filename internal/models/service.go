package models

import "github.com/google/uuid"

type Service struct {
	ServiceId   uuid.UUID `json:"service_id" db:"service_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Price       string    `json:"price" db:"price"`
	Category    string    `json:"category" db:"category"`
}
