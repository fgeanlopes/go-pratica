package dto

import "go-pratica/models"

// =====================================================
// DTO cliente
// Usado no endpoint: POST e PUT na api /api/v1/clients
// =====================================================
type ClientRequest struct {
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

// ToClientModel converte o DTO para o modelo de domínio
func (req *ClientRequest) ToClientModel() models.Client {
	return models.Client{
		Name:           req.Name,
		CPF:            req.CPF,
		PrimaryPhone:   req.PrimaryPhone,
		SecondaryPhone: req.SecondaryPhone,
		Email:          req.Email,
		ZipCode:        req.ZipCode,
		Street:         req.Street,
		Number:         req.Number,
		Complement:     req.Complement,
		Neighborhood:   req.Neighborhood,
		City:           req.City,
		State:          req.State,
		Status:         "active", // valor padrão
	}
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

func ToClientResponse(client models.Client) ClientResponse {
	return ClientResponse{
		ID:             client.ID,
		Name:           client.Name,
		CPF:            client.CPF,
		PrimaryPhone:   client.PrimaryPhone,
		SecondaryPhone: client.SecondaryPhone,
		Email:          client.Email,
		Status:         client.Status,
		CreatedAt:      client.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      client.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
