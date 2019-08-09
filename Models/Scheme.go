package Models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	FirstName    string
	LastName     string
	PhoneNumbers []PhoneNumber
}

type PhoneNumber struct {
	gorm.Model
	Number    string
	Name      string
	ContactId uint
}

func (c *Contact) TableName() string {
	return "contact"
}

func (p *PhoneNumber) TableName() string {
	return "phoneNumber"
}
