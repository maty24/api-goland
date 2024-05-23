package main

import (
	"api/config"
	"api/routes"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	// Configuración de la base de datos
	db := config.InitDB()

	// Obtener la conexión SQL subyacente
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error al obtener la conexión SQL: %v", err)
	}
	defer sqlDB.Close() // Cerrar la conexión al finalizar

	// Inicializar Echo
	e := echo.New()

	// Definir las rutas
	routes.InitRoutes(e, db)

	// Iniciar el servidor en el puerto 8080
	e.Logger.Fatal(e.Start(":8080"))
}
