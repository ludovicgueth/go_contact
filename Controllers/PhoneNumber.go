package Controllers

import (
	"api/ApiHelpers"
	"api/Models"
	"github.com/gin-gonic/gin"
)


func AddNewPhoneNumber(c *gin.Context) {
	var phoneNumber Models.PhoneNumber
	c.BindJSON(&phoneNumber)
	err := Models.AddNewPhoneNumberToContact(&phoneNumber)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, phoneNumber)
	} else {
		ApiHelpers.RespondJSON(c, 200, phoneNumber)
	}
}

func DeletePhoneNumber(c *gin.Context) {
	var phoneNumber Models.PhoneNumber
	id := c.Params.ByName("id")
	err := Models.DeletePhoneNumber(&phoneNumber, id)
	if err != nil {
		ApiHelpers.RespondJSON(c, 404, phoneNumber)
	} else {
		ApiHelpers.RespondJSON(c, 200, phoneNumber)
	}
}
