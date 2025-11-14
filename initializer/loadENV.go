package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv_variables() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error Loading ENV variables")
		return
	}
}
