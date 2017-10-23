package seed

import (
	"github.com/ethereal-go/example-call-center/root/database"
	"github.com/jinzhu/gorm"
)

type UsersSeed struct {

}

func (seed *UsersSeed) Run(db *gorm.DB) {
	user := &database.User{}

	user.Email = "test@sc.ru"
	user.Phone = "79299344312"
	user.Name = "Jo"

	db.NewRecord(user)
	db.Create(&user)
}