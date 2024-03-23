package models

import "time"

type Category struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	Category		string	  `json:"category"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
