package auth

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"

	"project-management-demo-backend/config"
	"project-management-demo-backend/pkg/entity/model"
)

type key string

const (
	// TokenKey is a key of user token.
	TokenKey key = "token"
)

// TokenID retrieve a token from context.
func TokenID(ctx context.Context) (string, error) {
	val := ctx.Value(TokenKey)

	t, ok := val.(*auth.Token)
	if !ok {
		return "", model.NewAuthError(errors.New("token is invalid type"))
	}

	tokenID := t.UID

	if tokenID == "" {
		return "", model.NewAuthError(errors.New("token is empty"))
	}

	return tokenID, nil
}

// New returns auth client
func New(ctx context.Context) (*auth.Client, error) {
	opt := option.WithCredentialsFile(config.C.Firebase.ServiceKey)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, model.NewInternalServerError(err)
	}

	a, err := app.Auth(ctx)
	if err != nil {
		return nil, model.NewAuthError(err)
	}

	return a, nil
}
