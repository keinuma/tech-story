package echo

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/keinuma/go-graphql/api/graph"
	"github.com/keinuma/go-graphql/api/graph/generated"
	"github.com/keinuma/go-graphql/api/infra/database/gorm"
)

func (s *Server) InitRouter(ctx context.Context) {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())

	conn := gorm.Connection

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{Resolvers: &graph.Resolver{
			DB: conn,
		}}))
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
