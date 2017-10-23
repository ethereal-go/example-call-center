package main

import (
	"github.com/ethereal-go/authJwtToken"
	"github.com/ethereal-go/base"
	"github.com/ethereal-go/ethereal"
	//"github.com/ethereal-go/example-call-center/root/database"
)

func main() {

	//db := ethereal.ConstructorDb()
	//db.DropTable(&database.Call{}, &database.User{})

	//db.Model(&database.Call{}).AddForeignKey("from_id", "users(id)", "CASCADE", "RESTRICT")
	//db.AutoMigrate(&database.Call{}, &database.User{})
	//db.Model(&database.Call{}).AddForeignKey("from_id", "users(id)", "RESTRICT", "RESTRICT")


	ethereal.Queries().Add("users", &base.UserField).Add("roles", &base.RoleField)
	ethereal.Mutations().Add("createUsers", &base.CreateUser).Add("token", &authJwtToken.CreateJWTToken)
	ethereal.ConstructorMiddleware().AddMiddleware(authJwtToken.GetMiddlewareJwtToken())
	authJwtToken.RegisterHandlerAuthCreateToken()

	ethereal.Start()
}
