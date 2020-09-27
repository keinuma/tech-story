package echo

import (
	"context"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/keinuma/tech-story/graphql"
	"github.com/keinuma/tech-story/graphql/generated"
	"github.com/keinuma/tech-story/infra/database/gorm"
	"github.com/keinuma/tech-story/infra/echo/auth"
)

func (s *Server) InitRouter(ctx context.Context) {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())
	s.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	conn := gorm.Connection
	if os.Getenv("APP_ENV") == "local" {
		conn = conn.Debug()
	}

	s.Engine.Use(auth.ForContext(ctx, conn))

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
