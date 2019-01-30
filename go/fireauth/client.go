package fireauth

import (
  "net/http"
  "os"
  "strings"

  "firebase.google.com/go/auth"
  "firebase.google.com/go"

  "catalyst-firewrap/firewrap"
  "github.com/Liquid-Labs/go-rest/rest"

  "golang.org/x/net/context"
)

type ScopedClient struct {
  client  *auth.Client
  request *http.Request
}

func GetClient(r *http.Request) (*ScopedClient, rest.RestError) {
  // TODO: verify that 'r.Context()' returns an app engine context
  // Initialize the app with a service account, granting admin privileges
	var app *firebase.App
	var err error
  var nodeEnv = os.Getenv("NODE_ENV")
	if nodeEnv != "production" || firewrap.Local {
		app, err = firebase.NewApp(r.Context(), firewrap.Config, firewrap.ClientOption)
	} else {
		app, err = firebase.NewApp(r.Context(), firewrap.Config)
	}
  if err != nil {
    return nil, rest.ServerError("Could not access authentication service.", err)
  }

  authClient, err := app.Auth(r.Context())
  if err != nil {
    return nil, rest.ServerError("Could not access authenticaiton service.", err)
  }

	return &ScopedClient{authClient, r}, nil
}

func (ab *ScopedClient) GetToken() (*auth.Token, rest.RestError) {
	authHeader := ab.request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	token, err := ab.client.VerifyIDToken(ab.Context(), tokenString)
	if err != nil {
		return nil, rest.AuthorizationError("Could not decode token.", err)
	}

	return token, nil
}

func (ab *ScopedClient) Context() (context.Context) {
  return ab.request.Context()
}
