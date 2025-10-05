package main

import (
	"gilangnyan/point-of-sales/internal/config"
	"gilangnyan/point-of-sales/package/database"
	"gilangnyan/point-of-sales/package/server"
	"log"
)

func main() {
	conf := config.Get()

	// Using interface-based database connection
	db, err := database.ConnectDatabase(database.PostgreSQL, *conf)
	if err != nil {
		log.Fatalf("Failed to initialize database: %s", err.Error())
	}

	// Ensure database connection is closed when main exits
	defer func() {
		if closeErr := db.Close(); closeErr != nil {
			log.Printf("Error closing database connection: %s", closeErr.Error())
		}
	}()

	server := server.NewGinServer(conf, db)
	server.Start()
}
