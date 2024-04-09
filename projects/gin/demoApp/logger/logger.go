package logger

import (
	"demoApp/config"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func Init() {
	customFormatter := new(logger.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"

	logger.SetFormatter(customFormatter)
	logger.SetReportCaller(true)

	logLevel := config.Appconfig.GetString("Logging.level")

	SetLogLevel(logLevel)

	if config.Appconfig.GetBool("Logging.stdout") {
		logger.New().Out = os.Stdout
	} else {
		file, err := os.OpenFile(config.Appconfig.GetString("Logging.path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		} else {
			fmt.Println("Failed to log to file", err.Error())
		}
	}
}

func SetLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	}

}

// logInfo
func LogInfo(msg string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
	}).Info(msg)
}

// logError
func LogError(msg string, err error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
		"error":        err.Error(),
	}).Error(msg)
}

// logFatal
func LogFatal(msg string, err error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("version"),
		"error":        err.Error(),
	}).Fatal(msg)
}

func LogDebug(msg, path, xRequestID string, errors error) {
	logger.WithFields(logger.Fields{
		"path":         path,
		"x-request-id": xRequestID,
		"version":      config.Appconfig.GetString("version"),
		"error":        "N/A",
	}).Debug(msg)
}

func PanicLn(msg string) {
	logger.Panicln(msg)
}

func FatalLn(msg string) {
	logger.Fatalln(msg)
}
func InfoLn(msg string) {
	logger.Infoln(msg)
}
func WarnLn(msg string) {
	logger.Warnln(msg)
}
func DebugLn(msg string) {
	logger.Debugln(msg)
}
func PrintLn(msg string) {
	logger.Println(msg)
}
