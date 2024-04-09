package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus
	router := gin.New()

	logrus.Println("Logrus")

	// SetReportCaller
	logrus.SetReportCaller(true)

	// SetFormatter
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		DisableTimestamp: false,
	})
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	// logrus SetOutput
	f, _ := os.Create("logrus.log")
	logrus.SetOutput(f)

	// logrus SetLevel
	logLevel()

	router.GET("/getdata", getData)
	router.Run(":3000")
}

func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "handler",
	})
}

func logLevel() {
	//logrus.SetLevel(logrus.TraceLevel)
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Traceln("trace")
	logrus.Debugln("debug")
	logrus.Infoln("info")
	logrus.Warnln("warn")
	logrus.Errorln("err")
	//logrus.Panicln("panic!!")
	//logrus.Fatalln("fatal")
}
