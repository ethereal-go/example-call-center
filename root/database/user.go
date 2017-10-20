package database

import (
	"github.com/ethereal-go/base/root/database"
)

type User struct {
	database.User
	Phone     string `json:"phone"`
}
