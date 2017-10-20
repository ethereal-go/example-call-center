package main

import (
	"github.com/ethereal-go/authJwtToken"
	"github.com/ethereal-go/base"
	"github.com/ethereal-go/ethereal"
)

func main() {

	ethereal.Queries().Add("users", &base.UserField).Add("roles", &base.RoleField)
	ethereal.Mutations().Add("createUsers", &base.CreateUser).Add("token", &authJwtToken.CreateJWTToken)
	ethereal.ConstructorMiddleware().AddMiddleware(authJwtToken.GetMiddlewareJwtToken())
	authJwtToken.RegisterHandlerAuthCreateToken()

	ethereal.Start()
}
