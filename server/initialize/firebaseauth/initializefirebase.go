package firebaseauth

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"log"
	"server/secrets/firebase"
	"strings"
)

var Client *auth.Client

func InitializeFirebaseAuth(PROJECT_ID string) {
	fb_cred := getFireBaseCred(PROJECT_ID)
	Client = createFirebaseClient(fb_cred)
}

func getFireBaseCred(PROJECT_ID string) string {
	fb_cred := ""
	if strings.Contains(PROJECT_ID, "local") {
		fb_cred = secrets_firebase.FireBaseLocal
	}
	if strings.Contains(PROJECT_ID, "stg") {
		fb_cred = secrets_firebase.FireBaseStg
	}
	if strings.Contains(PROJECT_ID, "prd") {
		fb_cred = secrets_firebase.FireBasePrd
	}
	log.Println("FireBase_SOURCE: ", fb_cred)
	if fb_cred == "" {
		log.Fatalln("no firebase source")
	}
	return fb_cred
}

func createFirebaseClient(CRED_PATH string) *auth.Client {
	opt := option.WithCredentialsFile(CRED_PATH)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	log.Println("\x1b[31m", "----- Firebase auth Initialize OK! -----", "\x1b[0m")

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	log.Println("\x1b[31m", "----- Firebase auth Client Spawn! -----", "\x1b[0m")
	return client
}
