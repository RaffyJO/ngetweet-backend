package main

import(
	"ngetweet/db/seeders"
	"ngetweet/db"
)

func main(){
	db.DatabaseInit()
	seeders.CreateCategory()
	seeders.CreateTweet()
	seeders.CreateUsers()
}