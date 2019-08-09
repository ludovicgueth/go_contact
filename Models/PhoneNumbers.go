package Models

import (
	"fmt"

	"api/Config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllPhoneNumber(b *[]PhoneNumber) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

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

func PutOnePhoneNumber(b *PhoneNumber, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeletePhoneNumber(b *PhoneNumber, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
