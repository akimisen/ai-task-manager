package model

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null;type:varchar(50)"`
	Email     string    `json:"email" gorm:"unique;not null;type:varchar(255)"`
	Password  string    `json:"-" gorm:"not null;type:varchar(255)"`
	Phone     string    `json:"phone" gorm:"type:varchar(20);unique;default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
