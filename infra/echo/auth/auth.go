package auth

import (
	"context"
	"errors"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"strings"

	"github.com/keinuma/tech-story/domain/service"
	"github.com/keinuma/tech-story/infra/database/gateway"
	"github.com/keinuma/tech-story/infra/firebase"
)

// Validator decodes the share session cookie and packs the session into context
func Validator(ctx context.Context, conn *gorm.DB) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get("Authorization")
			idToken := strings.Replace(authorization, "Bearer ", "", 1)

			if c.Path() == "/playground" {
				return next(c)
			}

			if idToken == "" {
				return errors.New("[auth.Validator] failed authentication, id token is empty")
			}
			uid, err := firebase.ValidateIDToken(idToken)
			if err != nil {
				return errors.New("[auth.Validator] failed authentication, failed validate token")
			}
			userService := service.NewUser(gateway.NewUser(ctx, conn))
			_, err = userService.GetUsersByUID(*uid)
			if err != nil {
				return errors.New("[auth.Validator] failed get user")
			}
			err = next(c)
			return err
		}
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context, conn *gorm.DB) echo.MiddlewareFunc {
	return Validator(ctx, conn)
}
