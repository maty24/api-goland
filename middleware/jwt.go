package middleware

import (
	"api/controllers"
	"github.com/labstack/echo/v4"
)

// Middleware para validar el token JWT
func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := validateToken(c)
			if err != nil { // Si hay un error, devolverlo
				c.Error(err)
				return err
			}
			return next(c)
		}
	}
}

// Función para validar el token JWT y me retorna un error si no es válido
func validateToken(c echo.Context) error {
	authController := controllers.NewAuthController()
	return authController.ValidateToken(c)
}
