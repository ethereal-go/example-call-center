package main

import  (
	"github.com/ethereal-go/authJwtToken"
	"github.com/ethereal-go/base"
	"github.com/ethereal-go/ethereal"
	"github.com/ethereal-go/example-call-center/database/seed"
)

func main() {
	userSeed := seed.UsersSeed{}
	userSeed.Run(ethereal.ConstructorDb())


	ethereal.Queries().Add("users", &base.UserField).Add("roles", &base.RoleField)
	ethereal.Mutations().Add("createUsers", &base.CreateUser).Add("token", &authJwtToken.CreateJWTToken)
	ethereal.ConstructorMiddleware().AddMiddleware(authJwtToken.GetMiddlewareJwtToken())
	authJwtToken.RegisterHandlerAuthCreateToken()
	ethereal.Start()
}
