// Routing (GET,POST)
// Handling Query String
// Handling URL params

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	router:=gin.Default()

	router.GET("/getdata",getData)
	router.POST("/postdata",postData)
	router.GET("/getquery",getQuery)
	router.GET("/geturlparam/:name/:age",getURLparam)

	router.Run(":3000")
	fmt.Println("server :3000")
}

func getData (ctx *gin.Context){
	data:=gin.H{
		"data":"Hi",
		"status":  http.StatusOK,
	}
	ctx.JSON(http.StatusOK,data)
}

func postData (ctx *gin.Context){
	body:=ctx.Request.Body
	value,_:=ioutil.ReadAll(body)

	data:=gin.H{
		"body":string(value),
	}

	ctx.JSON(200,data)
}

func getQuery(ctx *gin.Context){
//localhost:3000/getquery?name=Abc&age=18
    name:=ctx.Query("name") // ("key")
	age:=ctx.Query("age")

	data:=gin.H{
		"name":name,
		"age":age,
	}
	ctx.JSON(200,data)
}

func getURLparam(ctx *gin.Context){
	// localhost:3000/Xyz/20
	name:=ctx.Param("name")
	age:=ctx.Param("age")

	data:=gin.H{
		"name":name,
		"age":age,
	}
	ctx.JSON(200,data)
}
