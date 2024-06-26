package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/service/authservice"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"net/http"
	"strings"
)

func Auth(config authservice.Config, authInteractor *authservice.AuthInteractor) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := authInteractor.IsAuthorized(authToken, config.AccessTokenSecret)
				if authorized {
					claims, err := authInteractor.ParseAccessToken(authToken)
					if err != nil {
						return c.JSON(http.StatusUnauthorized, echo.Map{
							"status":  false,
							"message": message.ErrorMsg401UnAuthorized,
							"errors":  nil,
						})
					}
					claim.SetClaimsFromEchoContext(c, configs.AuthMiddlewareContextKey, claims)
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"status":  false,
					"message": message.ErrorMsg401UnAuthorized,
					"errors":  err.Error(),
				})
			}
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"status":  false,
				"message": message.ErrorMsg401UnAuthorized,
				"errors":  nil,
			})
		}
	}
}
