package models

import "time"

type Comment struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id"`
	TweetID         uint      `json:"tweet_id"`
	Body            string    `json:"body" gorm:"type:text"`
	ParentCommentID uint      `json:"parent_comment_id" gorm:"default:null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	ChildComment    []Comment `json:"child_comments" gorm:"foreignKey:ParentCommentID"`
}
