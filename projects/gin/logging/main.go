package main

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// logging

	router := gin.Default()

	// Write logging entries to file
	f, err := os.Create("logEntries.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	defer f.Close()

	router.GET("/getdata", getData)
	router.Run(":3000")

}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "getdata",
	})
}
