package firewrap

import (
	"log"
	"os"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

const apiKey string = "FIREBASE_API_KEY"
var Config *firebase.Config
var ClientOption option.ClientOption
var Local bool = false

func Setup(firebaseUrl string)  {
	Config = &firebase.Config{
		DatabaseURL: firebaseUrl,
	}
	// Fetch the service account key JSON file contents
	var envPurpose = os.Getenv(`NODE_ENV`)
	if envPurpose != "production" || Local {
		localFirebaseCredsFile := os.Getenv(apiKey)
		if localFirebaseCredsFile == "" {
			log.Panicf("'%s' required for local operation.", apiKey)
		}
		ClientOption = option.WithCredentialsFile(localFirebaseCredsFile)
	}
}
