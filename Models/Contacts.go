package Models

import (
	"fmt"

	"api/Config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetAllContact(b *[]Contact) (err error) {
	if err = Config.DB.Table("contacts").Select("firstname, lastname").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewContact(b *Contact) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneContact(b *Contact, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneContact(b *Contact, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteContact(b *Contact, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
