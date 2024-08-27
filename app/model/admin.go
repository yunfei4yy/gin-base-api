package model

import "time"

type Admin struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(64)" json:"name"`
	Password  string    `gorm:"column:password;type:varchar(64)" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
