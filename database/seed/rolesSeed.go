package seed

import (
	"github.com/ethereal-go/example-call-center/root/database"
	"github.com/jinzhu/gorm"
)

type RolesSeed struct{}

func (seed *RolesSeed) Run(db *gorm.DB) {

	var seeder = []*database.Role{
		{Name: "user", DisplayName: "This is user", Description: "Normal user"},
		{Name: "admin", DisplayName: "This is admin", Description: "Normal admin"},
	}

	for index := range seeder {
		role := database.Role{}
		db.Where("name = ?", seeder[index].Name).First(&role)
		if role.ID == 0 {
			db.Create(&seeder[index])
		}
	}
}
