package model

import "time"

type Account struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"not null"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	CurrencyID uint      `json:"currency_id" gorm:"not null"`
	Balance    uint      `json:"balance" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
