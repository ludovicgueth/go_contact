package Routers

import (
	"api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("contact", Controllers.ListContact)
		v1.POST("contact", Controllers.AddNewContact)
		v1.GET("contact/:id", Controllers.GetOneContact)
		v1.PUT("contact/:id", Controllers.PutOneContact)
		v1.DELETE("contact/:id", Controllers.DeleteContact)

		v1.POST("contact/:id/phoneNumber", Controllers.AddNewPhoneNumber)
		v1.DELETE("phoneNumber/:id", Controllers.DeletePhoneNumber)
		}

	return r
}
