package main

import (
	"log"

	"github.com/AdeleyeShina/go-crud/initializer"
	"github.com/AdeleyeShina/go-crud/models"
)

func init() {
	initializer.LoadEnv_variables()
	initializer.ConnectDB()
}

func main() {
	if err := initializer.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Error migrating")
		return
	}
}
