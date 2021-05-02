package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger
var outPath = "logs/log.txt"

func Default() *logrus.Logger {
	if Logger == nil {
		Logger = logrus.New()
		f, err := os.Create(outPath)

		if err != nil {
			Logger.Fatal(err)
		}
		// defer f.Close()
		Logger.SetOutput(f)
		Logger.SetLevel(logrus.InfoLevel)
		Logger.SetFormatter(&logrus.JSONFormatter{})
	}
	return Logger
}
