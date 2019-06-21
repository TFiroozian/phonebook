package logger

import (
	// Go native packages
	"os"

	// Dep packages
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	formatter := &logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		PrettyPrint:     true,
	}

	log := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: formatter,
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	log.SetReportCaller(true)
	return log
}
