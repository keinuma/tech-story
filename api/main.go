package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"github.com/keinuma/go-graphql/api/infra/database/gorm"
	"github.com/keinuma/go-graphql/api/infra/echo"
	"github.com/keinuma/go-graphql/api/infra/logger"
	"github.com/keinuma/go-graphql/api/library"
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
	gorm.InitDB()
	echo.Run()
}
