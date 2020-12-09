package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/keinuma/tech-story/infra/echo"
	"github.com/keinuma/tech-story/infra/firebase"
	"github.com/keinuma/tech-story/infra/logger"
	"github.com/keinuma/tech-story/library"
)

func init() {
	logger.Init()
	if library.IsLocal() {
		err := godotenv.Load()
		if err != nil {
			logrus.Error("Error loading .env file")
		}
	}
}

func main() {
	firebase.Init()
	echo.Run()
}
