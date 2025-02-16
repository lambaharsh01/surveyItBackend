package utils
import (
	"os"
	"log"
)

func GetEnv(key string) string {
	
	value := os.Getenv(key)
	var allowed bool = false

	if key == "DB_PASSWORD" {
		allowed = true
	}



	if value == "" && !allowed {
		log.Fatal("Error: " + key + "")
	}

	return value;
}