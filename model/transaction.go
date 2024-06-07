package model

import "time"

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null"`
	AccountID   uint      `json:"account_id" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	Category    Category  `json:"category" gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
