package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
	"os"
	secrets_cloudstorage "server/secrets/cloudstorage"
	"strings"
)

var Client *storage.Client
var Buckets secrets_cloudstorage.Buckets

func InitializeGcs(PROJECT_ID string) {
	setGcsCred(PROJECT_ID)
	Client = createGcsClient(PROJECT_ID)
	Buckets = getGcsBuckets(PROJECT_ID)
}

func setGcsCred(PROJECT_ID string) {
	secrets_json := ""
	if strings.Contains(PROJECT_ID, "local") {
		secrets_json = secrets_cloudstorage.GCS_Local
	} else if strings.Contains(PROJECT_ID, "prd") {
		secrets_json = secrets_cloudstorage.GCS_Prd
	} else if strings.Contains(PROJECT_ID, "stg") {
		secrets_json = secrets_cloudstorage.GCS_Stg
	} else {
		log.Fatalln("NO GCS Secrets")
	}
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", secrets_json)
	if err != nil {
		log.Fatalln(err)
	}
}

func createGcsClient(PROJECT_ID string) *storage.Client {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func getGcsBuckets(PROJECT_ID string) secrets_cloudstorage.Buckets {
	buckets := secrets_cloudstorage.Buckets{}
	if strings.Contains(PROJECT_ID, "local") {
		buckets = secrets_cloudstorage.Buckets_Local
	}
	if strings.Contains(PROJECT_ID, "stg") {
		buckets = secrets_cloudstorage.Buckets_Stg
	}
	if strings.Contains(PROJECT_ID, "prd") {
		buckets = secrets_cloudstorage.Buckets_Prd
	}
	log.Println("Buckets: ", buckets)
	return buckets
}
