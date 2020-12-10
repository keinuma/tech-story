package orm

import (
	"fmt"
	"gorm.io/gorm/logger"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	maxIdleConnsCount = 250
	maxOpenConnsCount = 125
)

func InitDB() *gorm.DB {
	dsn := GetConnectionString()
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
	}
	if os.Getenv("APP_ENV") == "local" {
		conn.Logger.LogMode(logger.Info)
		conn = conn.Debug()
	}
	return conn
}

func GetConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, dbname)
	logrus.Info(dsn)
	return dsn
}
