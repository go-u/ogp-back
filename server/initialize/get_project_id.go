package initialize

import (
	"log"
	"os"
)

func GetProjectId() string {
	PROJECT_ID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	log.Println("PROJECT_ID: ", PROJECT_ID)
	if PROJECT_ID == "" {
		log.Fatalln("Can't Get PROJECT_ID from env 'GOOGLE_CLOUD_PROJECT'\n If this is local test, Set 'appname-local' as GOOGLE_CLOUD_PROJECT")
	}
	return PROJECT_ID
}
