package controller

import (
	"demoApp/logger"
	"demoApp/model"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	model := model.GetData{
		Name: "Abc",
		Age:  22,
		City: "Bluru",
		Pin:  56001,
	}

	j, _ := json.Marshal(model)
	logger.LogInfo("In GetData", c)
	c.JSON(200, gin.H{
		"Data": string(j),
	})
}
