package models

import "time"

type Comment struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	TweetID         uint      `json:"tweet_id"`
	Body            string    `json:"body" gorm:"type:text"`
	ParentCommentID uint      `json:"parent_comment_id" gorm:"default:0"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
