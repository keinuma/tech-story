package echo

import (
	"context"
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
	s.InitRouter(ctx)
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
