package logger

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/keinuma/go-graphql/api/library"
)

func Init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if library.IsLocal() {
		logrus.SetOutput(os.Stdout)
	}
}
