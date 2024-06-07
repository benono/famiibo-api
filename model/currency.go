package model

import "time"

type Currency struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Code         string    `json:"code" gorm:"unique;not null"`
	JapaneseName string    `json:"japanese_name" gorm:"not null"`
	EnglishName  string    `json:"english_name" gorm:"not null"`
	Symbol       string    `json:"symbol" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
