package firewrap

import (
	"log"
	"os"

	"firebase.google.com/go"
	"google.golang.org/api/option"
)

const credsKey string = "CATALYST_FIREWRAP_FIREBASE_CREDS"
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
		localFirebaseCredsFile := os.Getenv(credsKey)
		if localFirebaseCredsFile == "" {
			log.Panicf("'%s' required for local operation.", credsKey)
		}
		ClientOption = option.WithCredentialsFile(localFirebaseCredsFile)
	}
}
