package seeders

import (
	"log"
	"ngetweet/db"
	"ngetweet/models"
)

func CreateUsers(){
	var data  = []models.User{
		{
			Name : "Raffy JO",
    Nickname : "JoQuery",
    Email : "Query@gmail.com",
    Avatar : ".jpg",
    Password : "rahasialah",
		},
		
		
}

for _,data1 := range data{
	db.DB.Create(&data1)
}

log.Println("seeded")
}