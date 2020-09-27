package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var AuthClient auth.Client

func Init() *auth.Client {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil
	}
	AuthClient = *authClient
	return authClient
}
