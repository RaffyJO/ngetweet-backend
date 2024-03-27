package models

import "time"

type User struct {
	ID              uint            `json:"id" gorm:"primaryKey"`
	Name            string          `json:"name" gorm:"not null;size:64"`
	Nickname        string          `json:"nickname" gorm:"unique;not null"`
	Email           string          `json:"email" gorm:"unique;not null"`
	Avatar          string          `json:"avatar" gorm:"default:null"`
	Password        string          `json:"password" gorm:"not null;size:64"`
	EmailVerifiedAt time.Time       `json:"email_verified_at" gorm:"default:null;type:datetime"`
	CreatedAt       time.Time       `json:"created_at" gorm:"type:datetime"`
	UpdatedAt       time.Time       `json:"updated_at" gorm:"type:datetime"`
	Tweets          []TweetResponse `json:"tweets" gorm:"foreignKey:UserID"`
	Likes           []Like          `json:"likes" gorm:"foreignKey:UserID"`
	Comments        []Comment       `json:"comments" gorm:"foreignKey:UserID"`
}
