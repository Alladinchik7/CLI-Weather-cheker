package system

import (
	"os"

	"github.com/joho/godotenv"
)

func LogApiKey() string {
	err := godotenv.Load("C:/Users/user/Desktop/PetProject/CLI_Go/.env")

	if err != nil {
		panic(err)
	}

	return os.Getenv("OPENWEATHER_API_KEY")
}