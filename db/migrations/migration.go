package migrations

import (
	"fmt"
	"ngetweet/db"
	"ngetweet/models"
)

func Migration() {
	// Drop existing tables
	// db.DB.Migrator().DropTable(&models.User{})
	// db.DB.Migrator().DropTable(&models.Tweet{})
	// db.DB.Migrator().DropTable(&models.Comment{})

	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Tweet{},
		&models.Comment{},
		&models.Like{},
	)

	if err != nil {
		fmt.Println("Can't running migrations")
	}

	fmt.Println("Migrated")
}
