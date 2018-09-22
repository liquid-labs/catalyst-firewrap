package firewrap

import (
	"firebase.google.com/go"

	"google.golang.org/appengine"
	"google.golang.org/api/option"

	"github.com/Liquid-Labs/go-webapp-tools/server"
)

var Config *firebase.Config
var ClientOption option.ClientOption

var Local bool = false

func Setup(firebaseUrl string)  {
	Config = &firebase.Config{
		DatabaseURL: firebaseUrl,
	}
	// Fetch the service account key JSON file contents
	if appengine.IsDevAppServer() || Local {
		localFirebaseCredsFile := server.MustGetenv("LOCAL_FIREBASE_CREDS")
		ClientOption = option.WithCredentialsFile(localFirebaseCredsFile)
	}
}
