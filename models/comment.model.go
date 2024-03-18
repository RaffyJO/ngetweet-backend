package models

import "time"

type Comment struct {
	ID              uint      	`json:"id" gorm:"primaryKey"`
	UserID          uint      	`json:"user_id"  `
	User			User	
	PostID          uint     	`json:"post_id"  `
	Tweet			Tweet		`gorm:"foreignKey:PostID"`
	Body            string   	`json:"body" gorm:"type:text"`
	ParentCommentID uint      	`json:"parent_comment_id"`
	ParentComment	[]Comment	`gorm:"foreignKey:PostID"`	
	CreatedAt       time.Time 	`json:"created_at"`
	UpdatedAt       time.Time 	`json:"updated_at"`
}
