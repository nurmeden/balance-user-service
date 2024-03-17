package models

import (
	"github.com/google/uuid"
	"time"
)

type RevenueReport struct {
	ReportID      uuid.UUID `json:"report_id" db:"report_id"`
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	RevenueAmount float64   `json:"revenue_amount" db:"revenue_amount"`
	ReportDate    time.Time `json:"report_date" db:"report_date"`
}
