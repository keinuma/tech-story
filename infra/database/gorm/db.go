package gorm

import (
	"fmt"
	"gorm.io/gorm/logger"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

const (
	maxIdleConnsCount = 250
	maxOpenConnsCount = 125
)

func InitDB() *gorm.DB {
	dsn := GetConnectionString(os.Getenv("DB_HOST"))
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	Connection = conn
	if os.Getenv("APP_ENV") == "local" {
		conn.Logger.LogMode(logger.Info)
	}
	return conn
}

func GetConnectionString(host string) string {
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	logrus.Info(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, dbname))
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		user, pass, host, port, dbname)
}
