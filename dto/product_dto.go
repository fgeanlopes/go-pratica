package dto

// =====================================================
// DTO para criar um novo cliente
// Usado no endpoint: POST /api/v1/clients
// =====================================================
type CreateClientRequest struct {
	Name           string `json:"name" binding:"required,min=3,max=255"`
	CPF            string `json:"cpf" binding:"required"`
	PrimaryPhone   string `json:"primary_phone" binding:"required"`
	SecondaryPhone string `json:"secondary_phone,omitempty"`
	Email          string `json:"email,omitempty" binding:"omitempty,email"`

	// Campos de endereço (opcionais na criação)
	ZipCode      string `json:"zip_code,omitempty"`
	Street       string `json:"street,omitempty"`
	Number       string `json:"number,omitempty"`
	Complement   string `json:"complement,omitempty"`
	Neighborhood string `json:"neighborhood,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty" binding:"omitempty,len=2"`
}

// =====================================================
// DTO para atualizar cliente existente
// Usado no endpoint: PUT /api/v1/clients/:id
// =====================================================
type UpdateClientRequest struct {
	Name           string `json:"name" binding:"omitempty,min=3,max=255"`
	PrimaryPhone   string `json:"primary_phone,omitempty"`
	SecondaryPhone string `json:"secondary_phone,omitempty"`
	Email          string `json:"email,omitempty" binding:"omitempty,email"`
	Status         string `json:"status,omitempty" binding:"omitempty,oneof=active inactive"`
}

// =====================================================
// DTO para resposta de cliente
// Usado em todas as respostas que retornam dados de cliente
// =====================================================
type ClientResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	CPF            string `json:"cpf"`
	PrimaryPhone   string `json:"primary_phone"`
	SecondaryPhone string `json:"secondary_phone,omitempty"`
	Email          string `json:"email,omitempty"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

// =====================================================
// DTO simplificado de cliente
// Usado quando cliente aparece dentro de outros objetos (OS, Veículo, etc)
// =====================================================
type ClientSimple struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	CPF          string `json:"cpf"`
	PrimaryPhone string `json:"primary_phone"`
}
