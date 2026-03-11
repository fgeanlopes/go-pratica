package controllers

import (
	"go-pratica/database"
	"go-pratica/dto"
	"go-pratica/models"
	"go-pratica/utils"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	// DTO para receber os dados da requisição
	var create dto.ClientRequest

	// Vincula o JSON da requisição ao DTO e valida os campos obrigatórios
	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// TODO: Validação redundante - ShouldBindJSON já valida usando as tags 'binding' do DTO.
	// Esta validação com govalidator pode ser removida no futuro, pois o DTO não possui tags 'valid:'.
	// Validação usando govalidator
	if _, err := govalidator.ValidateStruct(create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate := create.ToClientModel()

	cleanCpf, err := utils.ValidateCpf(clientCreate.CPF)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clientCreate.CPF = cleanCpf

	// Validação do CEP usando a função utilitária
	cleanZip, err := utils.ValidateZipCode(clientCreate.ZipCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate.ZipCode = cleanZip

	// Validação do telefone usando a função utilitária
	cleanCellPhone, err := utils.ValidateCellPhone(clientCreate.PrimaryPhone)

	// Se houver um erro na validação do telefone, retorna uma resposta de erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate.PrimaryPhone = cleanCellPhone

	// Validação do telefone secundário
	cleanTelePhone, err := utils.ValidateTelePhone(clientCreate.SecondaryPhone)

	// Se houver um erro na validação do telefone, retorna uma resposta de erro
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate.SecondaryPhone = cleanTelePhone

	// Criação do cliente no banco de dados
	if err := database.DB.Create(&clientCreate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente:" + err.Error()})
		return
	}

	// Retorna a resposta com os dados do cliente criado
	response := dto.ToClientResponse(clientCreate)
	c.JSON(http.StatusCreated, response)
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")

	clientID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var update dto.ClientRequest

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// TODO: Validação redundante - ShouldBindJSON já valida usando as tags 'binding' do DTO.
	// Esta validação com govalidator pode ser removida no futuro, pois o DTO não possui tags 'valid:'.
	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// GORM ignora automaticamente campos com zero-value (strings vazias, 0, false, etc.)
	// Atualiza apenas os campos que foram enviados na requisição
	updateData := update.ToClientModel()

	cleanCpf, err := utils.ValidateCpf(updateData.CPF)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData.CPF = cleanCpf

	// Validação do CEP usando a função utilitária
	cleanZip, err := utils.ValidateZipCode(updateData.ZipCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData.ZipCode = cleanZip

	// Validação do telefone primário
	cleanCellPhone, err := utils.ValidateCellPhone(updateData.PrimaryPhone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData.PrimaryPhone = cleanCellPhone

	// Validação do telefone secundário
	cleanTelePhone, err := utils.ValidateTelePhone(updateData.SecondaryPhone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateData.SecondaryPhone = cleanTelePhone

	result := database.DB.Model(&models.Client{}).Where("id = ?", uint(clientID)).Updates(updateData)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente:" + result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	// Busca o cliente atualizado para retornar na resposta
	var clientUpdate models.Client
	database.DB.First(&clientUpdate, uint(clientID))

	response := dto.ToClientResponse(clientUpdate)

	c.JSON(http.StatusOK, response)
}

func GetClients(c *gin.Context) {

	var getClients []models.Client

	if err := database.DB.Find(&getClients).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Clientes não encontrados"})
		return
	}

	var responses []dto.ClientResponse

	for _, p := range getClients {
		responses = append(responses, dto.ToClientResponse(p))
	}

	c.JSON(http.StatusOK, responses)
}

func GetClientByID(c *gin.Context) {
	id := c.Param("id")

	clientId, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var getClientId models.Client

	if err := database.DB.First(&getClientId, uint(clientId)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	response := dto.ToClientResponse(getClientId)

	c.JSON(http.StatusOK, response)
}

func DeleteClient(c *gin.Context) {
	id := c.Param("id")

	ClientId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var clientDelete models.Client

	result := database.DB.Delete(&clientDelete, uint(ClientId))

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente:" + result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
