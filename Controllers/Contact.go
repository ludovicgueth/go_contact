package Controllers

import (
	"api/ApiHelpers"
	"api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListContact(c *gin.Context) {
	// ONLY FIRSTNAME LASTNAME
	var contact []Models.Contact
	err := Models.GetAllContact(&contact)
	fmt.Printf("%+v\n", contact)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}

func AddNewContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	err := Models.AddNewContact(&contact)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}

func GetOneContact(c *gin.Context) {
	id := c.Params.ByName("id")
	var contact Models.Contact
	err := Models.GetOneContact(&contact, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}

func PutOneContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.GetOneContact(&contact, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	}
	c.BindJSON(&contact)
	err = Models.PutOneContact(&contact, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}

func DeleteContact(c *gin.Context) {
	var contact Models.Contact
	id := c.Params.ByName("id")
	err := Models.DeleteContact(&contact, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}
