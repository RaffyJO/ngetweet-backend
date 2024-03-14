package models

import "time"

type Tweet struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	UserID        uint           `json:"user_id"`
	Body          string         `json:"body" gorm:"type:text"`
	Image         string         `json:"image"`
	Likes         int            `json:"likes" gorm:"default:0"`
	TotalComments int            `json:"total_comments" gorm:"default:0"`
	CreatedAt     time.Time      `json:"created_at" gorm:"type:datetime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"type:datetime"`
	UserLikes     []LikeResponse `json:"user_likes" gorm:"foreignKey:TweetID"`
	Comments      []Comment      `json:"comments" gorm:"foreignKey:TweetID"`
}

type TweetResponse struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	UserID uint   `json:"-"`
	Body   string `json:"body" gorm:"type:text"`
	Image  string `json:"image"`
	Likes  int    `json:"likes" gorm:"default:0"`
}

func (TweetResponse) TableName() string {
	return "tweets"
}
