package models

import (
	"time"
)

type Like struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	UserID    uint `json:"user_id" gorm:"foreignKey:ID" josn:"user_id"`
	User      User
	TweetID   uint      `json:"tweet_id"`
	CreatedAt time.Time `json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:datetime"`
}

type LikeResponse struct {
	ID      uint `json:"id" gorm:"primaryKey"`
	UserID  uint `json:"user_id"`
	TweetID uint `json:"-"`
}

func (LikeResponse) TableName() string {
	return "likes"
}
