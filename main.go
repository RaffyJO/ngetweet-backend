package main

import (
	"ngetweet/db"
	"ngetweet/db/migrations"
	"ngetweet/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connection to database
	db.DatabaseInit()

	// Migration
	migrations.Migration()

	// Inisialisasi router Gin
	router := gin.Default()
	routes.RouteInit(router)
	router.Run(":8080")
}
