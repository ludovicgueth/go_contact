package main

import (
	"api/Config"
	"api/Models"
	"api/Routers"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("postgres", "host=db port=5432 user=user dbname=db password=password sslmode=disable")

	if err != nil {
		fmt.Println("statuse: ", err)
		log.Fatal("Error with DB")
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Contact{})
	Config.DB.AutoMigrate(&Models.PhoneNumber{})

	r := Routers.SetupRouter()
	// running
	r.Run()
}
