package models

import (
	"time"
)

type Followings struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID"`
	FollowingId uint      `json:"following_id"`
	Following   User      `gorm:"foreignKey:FollowingId"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:datetime"`
}
