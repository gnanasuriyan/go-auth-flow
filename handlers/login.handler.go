package handlers

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/wire"
	"oauth-study/models"
	"oauth-study/restapi/operations/user"
)

type ILoginHandler interface {
	Login(ctx context.Context, params user.LoginParams) middleware.Responder
}
type LoginHandler struct {
}

var NewLoginHandler = wire.NewSet(wire.Struct(new(LoginHandler), "*"), wire.Bind(new(ILoginHandler), new(LoginHandler)))

func (lh *LoginHandler) Login(ctx context.Context, params user.LoginParams) middleware.Responder {
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: "123123.12321.123123"})
}
