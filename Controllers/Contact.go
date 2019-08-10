package Controllers

import (
	"api/ApiHelpers"
	"api/Models"

	"github.com/gin-gonic/gin"
)

func ListContact(c *gin.Context) {
	var contact []Models.Contact
	err := Models.GetAllContact(&contact)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, contact)
	} else {
		ApiHelpers.RespondJSON(c, 200, contact)
	}
}

func AddNewContact(c *gin.Context) {
	var contact Models.Contact
	c.BindJSON(&contact)
	if contact.FirstName == "" {
		ApiHelpers.RespondJSON(c, 404, contact)
		return
	}
	if contact.LastName == "" {
		ApiHelpers.RespondJSON(c, 404, contact)
		return
	}
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
		var phoneNumbers []Models.PhoneNumber
		Models.GetPhoneNumbersFromContact(&phoneNumbers, id)
		contact.PhoneNumbers = phoneNumbers
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
