package controllers

import (
	"api/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ClienteController struct {
	DB *gorm.DB
}

type PedidoCliente struct {
	Nombre            string
	CorreoElectronico string
	FechaPedido       string
	Total             float64
}

func NewClienteController(DB *gorm.DB) *ClienteController {
	return &ClienteController{DB}
}

// Crear un nuevo cliente
func (cc *ClienteController) CreateCliente(c echo.Context) error {
	var cliente models.Cliente
	if err := c.Bind(&cliente); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result := cc.DB.Create(&cliente)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusCreated, cliente)
}

// Obtener todos los clientes
func (cc *ClienteController) GetClientes(c echo.Context) error {
	var clientes []models.Cliente
	cc.DB.Find(&clientes)

	// Verificar si hay clientes
	if len(clientes) == 0 {
		clientes = make([]models.Cliente, 0) // Crear un slice vacío
	}

	return c.JSON(http.StatusOK, clientes)
}

func (pc *ClienteController) GetPedidosWithClienteInfo(c echo.Context) error {
	var resultados []PedidoCliente

	// Ejecutar la consulta SQL
	pc.DB.Raw(`
		SELECT c.Nombre, c.Correo_Electronico, p.Fecha_Pedido, p.Total
		FROM Clientes AS c
		INNER JOIN Pedidos AS p ON c.ID_Cliente = p.ID_Cliente;
	`).Scan(&resultados)

	return c.JSON(http.StatusOK, resultados)
}

// Obtener un cliente por ID
func (cc *ClienteController) GetCliente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var cliente models.Cliente
	result := cc.DB.First(&cliente, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Cliente no encontrado")
	}

	return c.JSON(http.StatusOK, cliente)
}

// Actualizar un cliente
func (cc *ClienteController) UpdateCliente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var cliente models.Cliente
	if err := c.Bind(&cliente); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	cliente.ID_Cliente = uint(id) // Asegura que se actualice el cliente correcto

	result := cc.DB.Save(&cliente)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, cliente)
}

// Eliminar un cliente
func (cc *ClienteController) DeleteCliente(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var cliente models.Cliente
	result := cc.DB.Delete(&cliente, id)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
