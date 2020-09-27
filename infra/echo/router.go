package echo

import (
	"context"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"

	"github.com/keinuma/tech-story/graphql"
	"github.com/keinuma/tech-story/graphql/generated"
	"github.com/keinuma/tech-story/infra/database/gorm"
)

func (s *Server) InitRouter(ctx context.Context) {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())

	conn := gorm.Connection
	if os.Getenv("APP_ENV") == "local" {
		conn = conn.Debug()
	}
	logrus.Info(conn)

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &graphql.Resolver{
			DB: conn,
		}}))
	graphqlHandler.Use(extension.FixedComplexityLimit(7))

	playgroundHandler := playground.Handler("GraphQL", "/graphql")

	s.Engine.POST("/graphql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.Engine.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
