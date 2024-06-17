package model

import "time"

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount" gorm:"not null"`
	IsExpense   bool      `json:"is_expense" gorm:"not null"`
	Date        time.Time `json:"date" gorm:"not null"`
	AccountID   uint      `json:"account_id" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	CategoryID  uint      `json:"category_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	Amount        float64   `json:"amount" gorm:"not null"`
	IsExpense     bool      `json:"is_expense" gorm:"not null"`
	Date          time.Time `json:"date" gorm:"not null"`
	AccountID     uint      `json:"account_id" gorm:"not null"`
	AccountName   string    `json:"account_name"`
	PayeeID       uint      `json:"payee_id" gorm:"not null"`
	PayeeName     string    `json:"payee_name"`
	CategoryID    uint      `json:"category_id" gorm:"not null"`
	CategoryName  string    `json:"category_name"`
	CategoryIcon  string    `json:"category_icon"`
	CategoryColor string    `json:"category_color"`
	Description   string    `json:"description"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
