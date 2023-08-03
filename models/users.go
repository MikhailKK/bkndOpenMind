package models

import "time"

type User struct {
	// gorm.Model
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func (u User) GetName() string {
	return u.FirstName
}
