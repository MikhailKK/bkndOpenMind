package models

import "time"

type Users struct {
	// gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name" gorm:"not null"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	// RequestID uint   `json:"request_id" gorm:"foreignKey:RequestID"`
	// Request   Request `json:"request"`
}

func (u Users) GetName() string {
	return u.FirstName
}
