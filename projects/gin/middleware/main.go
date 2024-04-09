// Different ways to write middleware functions
// Apply middleware to routes , route groups , whole application
package main

import (
	"middleware/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// applyin MW to all routes
	router.Use(middleware.Authenticate)
	router.GET("/getdata", getData)
	router.GET("/getdata2", getData2)

	// applying to single route
	router.GET("/getdata", middleware.Authenticate, middleware.AddHeader, getData)

	//applying to routes group
	group := router.Group("/group", middleware.Authenticate)
	{
		group.GET("/getdata", getData)
		group.GET("/getdata2", getData2)
	}

	router.Run(":3000")

}

func getData(c *gin.Context) {
	data := gin.H{
		"msg": "getdata handler",
	}
	c.JSON(200, data)
}

func getData2(c *gin.Context) {
	data := gin.H{
		"msg": "getdata2 handler",
	}
	c.JSON(200, data)
}
