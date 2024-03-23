package models

import(
	"time"
)
type Followers struct{
	
	ID 				uint 		`json:"id" gorm:"primaryKey"` 
	UserID     		uint     	` json:"user_id"`
	User			User		`gorm:"foreignKey:UserID"`
	FollowersId	uint		`json:"following_id"`
	Followers		User		`gorm:"foreignKey:FollowersId"`
	CreatedAt     time.Time 	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt     time.Time 	`json:"updated_at" gorm:"type:datetime"`


}
