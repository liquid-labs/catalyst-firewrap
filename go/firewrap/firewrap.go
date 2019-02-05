package firewrap

import (
	"os"

	"firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/Liquid-Labs/go-api/osext"
)

const credsKey string = "FIREBASE_CREDS_FILE"
const dbUrlKey string = "FIREBASE_DB_URL"

var firebaseDbUrl = osext.MustGetenv(dbUrlKey)
var Config *firebase.Config
var ClientOption option.ClientOption
var Local bool = false

func Setup()  {
	Config = &firebase.Config{
		DatabaseURL: firebaseDbUrl,
	}
	// Fetch the service account key JSON file contents
	var envPurpose = os.Getenv(`NODE_ENV`)
	if envPurpose != "production" || Local {
		localFirebaseCredsFile := osext.MustGetenv(credsKey)
		ClientOption = option.WithCredentialsFile(localFirebaseCredsFile)
	}
}
