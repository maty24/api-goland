package models

import (
	"gorm.io/gorm"
	"time"
)

type Pedido struct {
	IDPedido    uint           `gorm:"primaryKey;column:id_pedido"` // Cambiamos a uint y agregamos la etiqueta column
	IDCliente   uint           `gorm:"not null;column:id_cliente"`
	Cliente     Cliente        `gorm:"foreignKey:IDCliente;references:ID_Cliente"` // Actualizamos la relación
	FechaPedido time.Time      `gorm:"not null;column:fecha_pedido"`
	Total       float64        `gorm:"not null;column:total"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
	UpdatedAt   time.Time      `gorm:"column:update_at"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
}
