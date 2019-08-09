package Models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	FirstName    string        `gorm:"not null"`
	LastName     string        `gorm:"not null"`
	PhoneNumbers []PhoneNumber `gorm:"foreignkey:ContactId"`
}

type PhoneNumber struct {
	gorm.Model
	Number    string
	Name      string
	ContactId uint
}
