package models

import (
	"time"

	"gorm.io/gorm"
)

// Vehicle - Modelo de veículo registrado na mecânica
type Vehicle struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	ClientID        uint   `gorm:"not null;index" json:"client_id"`
	Plate           string `gorm:"type:varchar(10);uniqueIndex;not null" json:"plate"`
	Brand           string `gorm:"type:varchar(100);not null" json:"brand"`
	Model           string `gorm:"type:varchar(100);not null" json:"model"`
	ManufactureYear int    `gorm:"not null" json:"manufacture_year"`
	ModelYear       int    `gorm:"not null" json:"model_year"`
	Color           string `gorm:"type:varchar(50)" json:"color,omitempty"`
	Chassis         string `gorm:"type:varchar(50)" json:"chassis,omitempty"`
	CurrentMileage  int    `gorm:"default:0" json:"current_mileage"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete

	// Relacionamentos
	Client Client `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE" json:"client,omitempty"`
	// ServiceOrders []ServiceOrder `gorm:"foreignKey:VehicleID;constraint:OnDelete:RESTRICT" json:"service_orders,omitempty"`
}

// TableName - Define o nome da tabela no banco
func (Vehicle) TableName() string {
	return "vehicles"
}
