package models

import (
	"time"

	"gorm.io/gorm"
)

// Client - Modelo de cliente da mecânica
type Client struct {
	ID             uint   `gorm:"primaryKey" json:"id"`
	Name           string `gorm:"type:varchar(255);not null" json:"name"`
	CPF            string `gorm:"type:varchar(14);uniqueIndex;not null" json:"cpf"`
	PrimaryPhone   string `gorm:"type:varchar(20);not null" json:"primary_phone"`
	SecondaryPhone string `gorm:"type:varchar(20)" json:"secondary_phone,omitempty"`
	Email          string `gorm:"type:varchar(255)" json:"email,omitempty"`
	Status         string `gorm:"type:enum('active','inactive');default:'active'" json:"status"`

	ZipCode      string `gorm:"type:varchar(10)" json:"zip_code,omitempty"`
	Street       string `gorm:"type:varchar(255)" json:"street,omitempty"`
	Number       string `gorm:"type:varchar(20)" json:"number,omitempty"`
	Complement   string `gorm:"type:varchar(255)" json:"complement,omitempty"`
	Neighborhood string `gorm:"type:varchar(100)" json:"neighborhood,omitempty"`
	City         string `gorm:"type:varchar(100)" json:"city,omitempty"`
	State        string `gorm:"type:char(2)" json:"state,omitempty"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Soft delete

	// Relacionamentos
	// Addresses []Address `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE" json:"addresses,omitempty"`
	// Vehicles  []Vehicle `gorm:"foreignKey:ClientID;constraint:OnDelete:CASCADE" json:"vehicles,omitempty"`
}

// TableName - Define o nome da tabela no banco
func (Client) TableName() string {
	return "clients"
}
