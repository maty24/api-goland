package controllers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

type AuthController struct{} // No necesitas una instancia de la base de datos

// esto es un constructor que retorna una instancia de AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// esto es un metodo que genera un token
func (ac *AuthController) GenerateToken(c echo.Context) error {
	// Crea el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Expira en 24 horas
	})

	// Firma el token con tu clave secreta (reemplaza "secreto" por una clave real)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error al generar el token"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}

// funcion que valida el token
func (ac *AuthController) ValidateToken(c echo.Context) error {
	// Obtener el token de la cabecera Authorization
	authHeader := c.Request().Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Parsear y validar el token
	err := ac.validateJWTToken(tokenString)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "EHH Maquinola quien chota sos?, token inválido")
	}
	return nil // Token válido
}

// funcion que valida el token
func (ac *AuthController) validateJWTToken(tokenString string) error {
	//el jwt.parse recibe el token y una funcion que valida el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Asegúrate de que el método de firma es el correcto (HS256 en este caso)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(http.StatusUnauthorized, "Método de firma inválido")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return err
	}

	return nil
}
