package seeders

import (
	"log"
	"ngetweet/db"
	"ngetweet/models"
)

func CreateCategory(){
	var data  = []models.Category{
				{ID: 1,Category : "Kimia Kelas 10 SMA"},
				{ID: 2,Category: "Fisika Kelas 10 SMA"},
				{ID: 3,Category: "Biologi Kelas 10 SMA"},
				{ID: 4,Category: "Matematika Peminatan Kelas 10 SMA"},
				{ID: 5,Category: "Matematika Wajib Kelas 10 SMA"},
				{ID: 6,Category: "Bahasa Indonesia Kelas 10 SMA"},
				{ID: 7,Category: "PPKn Kelas 10 SMA"},
				{ID: 8,Category: "Bahasa Inggris Kelas 10 SMA"},
				{ID: 9,Category: "Sejarah Indonesia Kelas 10 SMA"},
				{ID: 10,Category: "Ekonomi Kelas 10 SMA"},
	}

	for _,data1 := range data{
		db.DB.Create(&data1)
		// if result.Error!=nil{
		// 	log.Println(result.Error.Error())
		// }
	}

	log.Println("Seeded")
	
}