// custom HTTP config
// Route Grouping
// Basic Auth

package main

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// basic Auth
	accounts := gin.Accounts{
		"user": "pass123",
		"me":   "getIN",
	}
	auth := gin.BasicAuth(accounts)

	// Route Grouping
	admin := router.Group("/admin", auth) // auth is reqd. for all admin routes
	{
		admin.GET("/getdata", getData)
		admin.GET("/getquery", getQuery)
	}

	client := router.Group("/client")
	{
		client.POST("/postdata", postData)
	}

	// custom HTTP config
	server := &http.Server{
		Addr:         ":3000", // port
		Handler:      router,  // every request is passed on to router
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe() // router.Run(":3000")
}

func getData(ctx *gin.Context) {
	data := gin.H{
		"data":   "Hi",
		"status": http.StatusOK,
	}
	ctx.JSON(http.StatusOK, data)
}

func postData(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := io.ReadAll(body)
	data := gin.H{
		"body": string(value),
	}
	ctx.JSON(200, data)
}

func getQuery(ctx *gin.Context) {
	name := ctx.Query("name")
	age := ctx.Query("age")
	data := gin.H{
		"name": name,
		"age":  age,
	}
	ctx.JSON(200, data)
}
