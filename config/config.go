package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv()  {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
}

/* GetEnv : get environment variable or return default value
   key : nama environment variable
   defaultValue : nilai default jika variable tidak ditemukan
   return : nilai environment variable atau default value
*/
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return defaultValue
	}

	return value
}