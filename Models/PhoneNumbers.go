package Models

import (
	"api/Config"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AddNewPhoneNumber(b *PhoneNumber) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOnePhoneNumber(b *PhoneNumber, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func DeletePhoneNumber(b *PhoneNumber, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
