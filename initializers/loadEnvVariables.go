package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVairables() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loding .env file")
	}
}