package middleware

import (
	"api/controllers"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Middleware para validar el token JWT
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authController := controllers.NewAuthController()

			err := authController.ValidateToken(c)
			if err != nil {
				c.Error(err)
				// Devuelve un error HTTP 401 para indicar fallo en la autenticación
				return echo.NewHTTPError(http.StatusUnauthorized, "No autorizado")
			}

			return next(c)
		}
	}
}
