package fireauth

import (
  "fmt"
  "strings"

  "firebase.google.com/go/auth"

  "github.com/Liquid-Labs/terror/go/terror"
)

func (ab *ScopedClient) CheckAuthorizedAll(reqClaims ...string) (*auth.Token, terror.Terror) {
	token, err := ab.GetToken()
	if err != nil {
		return nil, err
	}

	claims := token.Claims
	for _, reqClaim := range reqClaims {
		claim, ok := claims[reqClaim]
		if !ok || !claim.(bool) {
			return nil, terror.AuthorizationError(fmt.Sprintf("User '%s' failed to access resource requiring claim '%s'.", token.UID, reqClaim), nil)
		}
	}

	return token, nil
}

func (ab *ScopedClient) CheckAuthorizedAny(reqClaims ...string) (*auth.Token, terror.Terror) {
	token, err := ab.GetToken()
	if err != nil {
		return nil, err
	}

	if len(reqClaims) == 0 {
		return token, nil
	}

	claims := token.Claims
	for _, reqClaim := range reqClaims {
		claim, ok := claims[reqClaim]
		if ok && claim.(bool) {
			return token, nil
		}
	}

	return nil, terror.AuthorizationError(
		fmt.Sprintf("User '%s' failed to access resource requiring at least one claim '%s'.",
			 token.UID, strings.Join(reqClaims, ", ")), nil)
}
