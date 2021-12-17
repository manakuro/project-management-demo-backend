package auth

import (
	"context"
	"errors"
	"project-management-demo-backend/pkg/entity/model"

	"firebase.google.com/go/auth"
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
