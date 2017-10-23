package seed

import (
	"github.com/ethereal-go/example-call-center/root/database"
	"github.com/jinzhu/gorm"
)

type RolesSeed struct {}

func (seed *RolesSeed) Run(db *gorm.DB) {
	role := database.Role{}

	role.Name = "user"
	role.DisplayName = "This is user"


	//db.NewRecord(user)
	db.Create(&role)
}
