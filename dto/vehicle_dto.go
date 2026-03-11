package dto

type VehicleRequest struct {
	ClientID uint `json:"client_id" binding:"required,min=1"`
	//TODO Validar a placa com regex para aceitar os formatos brasileiros (ABC-1234 ou ABC1D23)
	// Colocar regra no model para validar a placa
	Plate           string `json:"plate" binding:"required,min=7,max=8"`
	Brand           string `json:"brand" binding:"required,min=2,max=100"`
	Model           string `json:"model" binding:"required,min=2,max=100"`
	ManufactureYear int    `json:"manufacture_year" binding:"required,min=1886"`
	Color           string `json:"color" binding:"required,min=3,max=50"`
	Chassis         string `json:"chassis" binding:"required,min=17,max=17"`
	CurrentMileage  int    `json:"current_mileage" binding:"required,min=1"`
	//TODO O ano mínimo é 1886, quando o primeiro carro foi inventado,
	// e o máximo é o ano atual + 1 para permitir modelos do próximo ano
	// Colocar regra no model para validar o ano
	ModelYear int `json:"model_year" binding:"required,min=1886"`
}

// func (req *VehicleRequest) ToVehicleModel() models.Vehicle {
// 	return models.vehicle{}
// }
