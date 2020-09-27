package firebase

import (
	"context"
)

func ValidateIDToken(idToken string) (*string, error) {
	token, err := AuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return &token.UID, nil
}
