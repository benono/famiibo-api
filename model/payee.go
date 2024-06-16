package model

import "time"

type Payee struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FamilyID  uint      `json:"family_id" gorm:"not null;uniqueIndex:idx_family_name"`
	Name      string    `json:"name" gorm:"not null;uniqueIndex:idx_family_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
