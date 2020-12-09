package echo

import (
	"context"
	"github.com/keinuma/tech-story/infra/database/orm"
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
	s.InitRouter(ctx, conn)
	s.Engine.HideBanner = true
	s.Engine.HidePort = true
	logrus.Debug("starting api server")
	s.Engine.Start(":" + port())
}

func port() string {
	defaultPort := "8000"
	envPort := os.Getenv("PORT")
	if envPort == "" {
		return defaultPort
	}
	return envPort
}
