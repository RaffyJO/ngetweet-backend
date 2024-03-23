package seeders

import (
	"log"
	"ngetweet/db"
	"ngetweet/models"
)

func CreateTweet(){
	var data  = []models.Tweet{
		{
			UserID:1,
			Body:"Apa fungsi alat kelamin ?",
			CategoryId:3,
		},
		{
			UserID:1,
			Body:"Apa rumus gaya ?",
			CategoryId:2,
		},
		{
			UserID:1,
			Body:"Jelaskan sejarah rengasdengklok ?",
			CategoryId:10,
		},
		
}

for _,data1 := range data{
	db.DB.Create(&data1)
}

log.Println("seeded")
}