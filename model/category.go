package model

import "time"

type Category struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Order     uint      `json:"order" gorm:"not null"`
	Name      string    `json:"name" gorm:"not null"`
	FamilyID  uint      `json:"family_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
