package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ClienteRoutes(e *echo.Echo, db *gorm.DB) {
	// Crear instancias de los controladores, es como llamar la clase por asi decirlo
	clienteController := controllers.NewClienteController(db)

	// Preparar grupos de rutas
	clientesGroup := e.Group("/clientes")

	// Rutas para clientes
	clientesGroup.POST("", clienteController.CreateCliente, middleware.JWTMiddleware())
	clientesGroup.GET("", clienteController.GetClientes, middleware.JWTMiddleware())
	clientesGroup.GET("/sql", clienteController.GetPedidosWithClienteInfo)
	clientesGroup.GET("/:id", clienteController.GetCliente)
	clientesGroup.PUT("/:id", clienteController.UpdateCliente)
	clientesGroup.DELETE("/:id", clienteController.DeleteCliente)
}

func PedidoRoutes(e *echo.Echo, db *gorm.DB) {
	pedidoController := controllers.NewPedidoController(db)
	pedidosGroup := e.Group("/pedidos")
	pedidosGroup.POST("", pedidoController.CreatePedido)
	pedidosGroup.GET("", pedidoController.GetPedidos)
	pedidosGroup.GET("/:id", pedidoController.GetPedido)
}

func AuthRoutes(e *echo.Echo) {
	authController := controllers.NewAuthController()
	authGroup := e.Group("/auth")
	authGroup.GET("/token", authController.GenerateToken)
}
