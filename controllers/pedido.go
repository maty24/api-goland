package controllers

import (
	"api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PedidoController struct {
	DB *gorm.DB
}

func NewPedidoController(DB *gorm.DB) *PedidoController {
	return &PedidoController{DB}
}

// Crear un nuevo pedido
func (pc *PedidoController) CreatePedido(c echo.Context) error {
	var pedido models.Pedido
	if err := c.Bind(&pedido); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validar que el IDCliente exista
	var cliente models.Cliente
	if err := pc.DB.First(&cliente, pedido.IDCliente).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Cliente no encontrado")
	}

	result := pc.DB.Create(&pedido)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, pedido)
}

// Obtener todos los pedidos con sus clientes asociados
func (pc *PedidoController) GetPedidos(c echo.Context) error {
	var pedidos []models.Pedido
	pc.DB.Preload("Cliente").Find(&pedidos) // El preload carga los datos del cliente que está asociado al pedido, la relacion debe estar definida en el modelo
	return c.JSON(http.StatusOK, pedidos)
}

// Obtener todos los pedidos sin los datos del cliente asociado
func (pc *PedidoController) GetPedidosWithoutClient(c echo.Context) error {
	var pedidos []models.Pedido
	pc.DB.Find(&pedidos) // No cargar los datos del cliente
	return c.JSON(http.StatusOK, pedidos)
}

// Obtener un pedido por ID con su cliente asociado
func (pc *PedidoController) GetPedido(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var pedido models.Pedido
	result := pc.DB.Preload("Cliente").First(&pedido, id) // Cargar el cliente
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Pedido no encontrado")
	}

	return c.JSON(http.StatusOK, pedido)
}
