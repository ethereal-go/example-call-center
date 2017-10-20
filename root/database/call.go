package database

import "time"

type Call struct {
	ID                  uint   `json:"id";gorm:"primary_key"`
	From                User   `json:"from"`
	FromID              int    `gorm:"index"`
	To                  User   `json:"to"`
	ToID                int    `gorm:"index"`
	OutgoingIncoming    string `json:"outgoing_incoming"`
	CoordinatesIncoming string `json:"coordinates_incoming"`
	Duration            int    `json:"duration"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *time.Time `sql:"index"`
}
