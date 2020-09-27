package firebase

import (
	"context"
	"github.com/sirupsen/logrus"
)

func ValidateIDToken(idToken string) (*string, error) {
	token, err := AuthClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &token.UID, nil
}
