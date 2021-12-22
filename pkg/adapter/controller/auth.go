package controller

import (
	"net/http"
	"project-management-demo-backend/pkg/adapter/handler"
	"project-management-demo-backend/pkg/entity/model"
	uauth "project-management-demo-backend/pkg/util/auth"
)

type auth struct{}

// Auth is an interface of controller
type Auth interface {
	RevokeRefreshTokens(ctx Context) error
}

// NewAuthController generates test auth controller
func NewAuthController() Auth {
	return &auth{}
}

func (a *auth) RevokeRefreshTokens(ctx Context) error {
	c := ctx.Request().Context()

	var params struct {
		UID string `json:"uid"`
	}
	if err := ctx.Bind(&params); err != nil {
		return handler.HandleRestError(ctx, model.NewInvalidParamError(err, nil))
	}

	client, err := uauth.New(c)
	if err != nil {
		return err
	}

	if err = client.RevokeRefreshTokens(c, params.UID); err != nil {
		return handler.HandleRestError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, nil)
}
