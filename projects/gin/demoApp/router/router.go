package router

import (
	"demoApp/config"
	"demoApp/controller"
	"demoApp/middleware"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := Newrouter()
	router.Run(config.Appconfig.GetString("server.port"))
}

func Newrouter() *gin.Engine {
	router := gin.New()
	resource := router.Group("/api")
	resource.Use(middleware.LogRequestInfo())
	{
		resource.GET("/GetData", controller.GetData)
	}
	return router
}
