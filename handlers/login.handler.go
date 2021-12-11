package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"oauth-study/models"
	"oauth-study/restapi/operations/user"
)

type LoginHandler struct {
}

func (lh *LoginHandler) Handle(params user.LoginParams) middleware.Responder {
	return user.NewLoginOK().WithPayload(&models.LoginSuccess{Success: true, Token: "123123.12321.123123"})
}
