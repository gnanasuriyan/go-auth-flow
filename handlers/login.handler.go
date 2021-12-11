package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"oauth-study/models"
	"oauth-study/restapi/operations/user"
)

type ILoginHander interface {
	Login(ctx context.Context, params user.LoginParams) middleware.Responder
}
type LoginHandler struct {
}

func (lh *LoginHandler) Login(ctx context.Context, params user.LoginParams) middleware.Responder {
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: "123123.12321.123123"})
}
