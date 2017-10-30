package seed

import (
	"github.com/ethereal-go/example-call-center/root/database"
	"github.com/jinzhu/gorm"
)

type UsersSeed struct{}

func (seed *UsersSeed) Run(db *gorm.DB) {

	var seeder = []*database.User{
		{
			Email: "a@flskd.com",
			Phone: "a@flskd.com",
			Name:  "Agaria",
		},
	}
	for index := range seeder {
		user := database.User{}
		db.Where("name = ?", seeder[index].Name).First(&role)
		if user.ID == 0 {
			db.Create(&seeder[index])
		}
	}
}
