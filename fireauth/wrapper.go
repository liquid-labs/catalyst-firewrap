package fireauth

import (
  "firebase.google.com/go/auth"

  "github.com/Liquid-Labs/go-webapp-tools/rest"
)

func (ab *ScopedClient) GetUser(subject string) (*auth.UserRecord, rest.RestError) {
	userRecord, err := ab.client.GetUser(ab.Context(), subject)
	if err != nil {
		return nil, rest.ServerError("Could not retrieve user record.", err)
	}

	return userRecord, nil
}

func (ab *ScopedClient) Users(nextPageToken string) (*auth.UserIterator) {
  return ab.client.Users(ab.Context(), nextPageToken)
}
