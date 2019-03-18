package impl

import (
	// Dep packages
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	contact := router.Group("/api/v0/contacts")
	{
		contact.GET(":id", GetContact)
		contact.GET("", ListContact)
		contact.DELETE(":id", DeleteContact)
		contact.POST("", CreateContact)
	}

	return router
}
