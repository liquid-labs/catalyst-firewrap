package fireauth

import (
  "firebase.google.com/go/auth"

  "github.com/Liquid-Labs/terror/go/terror"
)

func (ab *ScopedClient) GetUser(subject string) (*auth.UserRecord, terror.Terror) {
	userRecord, err := ab.client.GetUser(ab.Context(), subject)
	if err != nil {
		return nil, terror.ServerError("Could not retrieve user record.", err)
	}

	return userRecord, nil
}

func (ab *ScopedClient) Users(nextPageToken string) (*auth.UserIterator) {
  return ab.client.Users(ab.Context(), nextPageToken)
}
