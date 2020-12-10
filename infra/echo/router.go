package echo

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"

	"github.com/keinuma/tech-story/graphql"
	"github.com/keinuma/tech-story/graphql/generated"
	"github.com/keinuma/tech-story/infra/echo/auth"
)

func (s *Server) InitRouter(ctx context.Context, conn *gorm.DB, storeConn *redis.Conn) {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())
	s.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	s.Engine.Use(auth.ForContext(ctx, conn))

	graphqlHandler := handler.NewDefaultServer(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graphql.Resolver{
				DB:        conn,
				StorePool: storeConn,
			},
		}),
	)
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
