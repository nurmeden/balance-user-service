package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	UserID        uuid.UUID `json:"user_id" db:"user_id"`
	FirstName     string    `json:"first_name" db:"user_id"`
	LastName      string    `json:"last_name" db:"user_id"`
	Email         string    `json:"email,omitempty" db:"user_id"`
	Password      string    `json:"password,omitempty" db:"user_id"`
	Gender        string    `json:"gender" db:"gender"`
	Balance       int64     `json:"balance" db:"user_id"`
	Role          string    `json:"role" db:"user_id"`
	City          string    `json:"city" db:"city"`
	Country       string    `json:"country" db:"country"`
	Address       string    `json:"address" db:"user_id"`
	Status        string    `json:"status" db:"user_id"`
	CreatedAt     time.Time `json:"created_at" db:"user_id"`
	LastLoginDate time.Time `json:"last_login_date" db:"user_id"`
}

func (u *User) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (u *User) ComparePasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) GetDisplayName() string {
	return u.FirstName + u.LastName
}
