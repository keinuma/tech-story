package echo

import (
	"context"
	"github.com/keinuma/tech-story/infra/database/orm"
	"github.com/keinuma/tech-story/infra/store"
	"os"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Engine *echo.Echo
}

func newServer() *Server {
	return &Server{
		Engine: echo.New(),
	}
}

func Run() {
	ctx := context.Background()
	s := newServer()
	conn := orm.InitDB()
	storeConn, err := store.NewRedisClient()
	if err != nil {
		os.Exit(1)
	}
	s.InitRouter(ctx, conn, storeConn)
	s.Engine.HideBanner = true
	s.Engine.HidePort = true
	logrus.Debug("starting api server")
	err = s.Engine.Start(":" + port())
	if err != nil {
		os.Exit(1)
	}
}

func port() string {
	defaultPort := "8000"
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return defaultPort
	}
	return envPort
}
