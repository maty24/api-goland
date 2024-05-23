package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Config struct {
	DB_Host     string
	DB_User     string
	DB_Password string
	DB_Name     string
	DB_Port     string
}

func InitDB() *gorm.DB {
	// Obtén los valores de configuración del entorno (Docker)
	config := Config{
		DB_Host:     "10.6.22.9", // Nombre del servicio en docker-compose
		DB_User:     "postgres",
		DB_Password: "1234",
		DB_Name:     "goland",
		DB_Port:     "5432",
	}

	// Construye la cadena de conexión DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Santiago",
		config.DB_Host, config.DB_User, config.DB_Password, config.DB_Name, config.DB_Port)

	// Establece la conexión a la base de datos
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}

	// Retorna la conexión
	return db
}
