package routes

import (
	"api/controllers"
	"api/middleware"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	// Crear instancias de los controladores
	clienteController := controllers.NewClienteController(db)
	pedidoController := controllers.NewPedidoController(db)
	authController := controllers.NewAuthController()

	// Preparar grupos de rutas
	clientesGroup := e.Group("/clientes")
	pedidosGroup := e.Group("/pedidos")

	// Rutas para clientes
	clientesGroup.POST("", clienteController.CreateCliente)
	clientesGroup.GET("", clienteController.GetClientes, middleware.JWTMiddleware())
	clientesGroup.GET("/sql", clienteController.GetPedidosWithClienteInfo, middleware.JWTMiddleware())
	clientesGroup.GET("/:id", clienteController.GetCliente)
	clientesGroup.PUT("/:id", clienteController.UpdateCliente)
	clientesGroup.DELETE("/:id", clienteController.DeleteCliente)

	// Rutas para pedidos
	pedidosGroup.POST("", pedidoController.CreatePedido)
	pedidosGroup.GET("", pedidoController.GetPedidos, middleware.JWTMiddleware())
	pedidosGroup.GET("/:id", pedidoController.GetPedido)
	// ... (Rutas para actualizar y eliminar pedidos)

	e.GET("/token", authController.GenerateToken)
}
