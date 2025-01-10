package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
	Log.SetFormatter(&logrus.TextFormatter{})
}
