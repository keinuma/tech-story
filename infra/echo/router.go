package echo

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/keinuma/tech-story/graph"
	"github.com/keinuma/tech-story/graph/generated"
	"github.com/keinuma/tech-story/infra/echo/auth"
	"github.com/keinuma/tech-story/infra/store"
)

func (s *Server) InitRouter(ctx context.Context, conn *gorm.DB, storeConn *store.Store, subscriber *store.Subscriber) {
	s.Engine.Use(middleware.Logger())
	s.Engine.Use(middleware.Recover())
	s.Engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	s.Engine.Use(auth.ForContext(ctx, conn))
	graphqlHandler := handler.New(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: &graph.Resolver{
				DB:         conn,
				StorePool:  storeConn,
				Subscriber: subscriber,
			},
		}),
	)
	graphqlHandler.AddTransport(transport.Options{})
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	graphqlHandler.Use(extension.FixedComplexityLimit(7))

	playgroundHandler := playground.Handler("GraphQL", "/gql")

	s.Engine.POST("/gql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.Engine.GET("/gql", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	s.Engine.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
