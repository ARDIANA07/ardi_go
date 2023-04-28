package entity

import "time"

type User struct {
	ID             int
	Name           string `gorm:"type:varchar(255);not null"`
	Occupation     string `gorm:"type:varchar(255);not null"`
	Email          string `gorm:"type:varchar(255);not null"`
	PasswordHas    string `gorm:"type:varchar(255);not null"`
	AvatarFileName string `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time
	UpdateAt       time.Time
}
