package controllers

import (
	"go-pratica/database"
	"go-pratica/dto"
	"go-pratica/models"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateClient(c *gin.Context) {
	var create dto.ClientRequest

	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// Validação usando govalidator
	if _, err := govalidator.ValidateStruct(create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate := create.ToClientModel()

	if err := database.DB.Create(&clientCreate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente:" + err.Error()})
		return
	}

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

	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var clientUpdate models.Client

	if err := database.DB.First(&clientUpdate, uint(clientID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	// GORM ignora automaticamente campos com zero-value (strings vazias, 0, false, etc.)
	// Atualiza apenas os campos que foram enviados na requisição
	updateData := update.ToClientModel()

	if err := database.DB.Model(&clientUpdate).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar cliente:" + err.Error()})
		return
	}

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

	// TODO otimizar para evitar 2 round-trips
	if err := database.DB.First(&clientDelete, uint(ClientId)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}

	if err := database.DB.Delete(&clientDelete, uint(ClientId)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar cliente:" + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
