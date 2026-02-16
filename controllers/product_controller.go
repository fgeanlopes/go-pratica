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
	var create dto.CreateClientRequest

	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// Validação usando govalidator
	if _, err := govalidator.ValidateStruct(create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientCreate := models.Client{
		Name:  create.Name,
		Price: create.Price,
	}

	if err := database.DB.Create(&clientCreate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto:" + err.Error()})
		return
	}

	response := dto.ClientResponse{
		ID:        clientCreate.ID,
		Name:      clientCreate.Name,
		Price:     clientCreate.Price,
		CreatedAt: clientCreate.CreatedAt,
		UpdatedAt: clientCreate.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func UpdateClient(c *gin.Context) {
	id := c.Param("id")

	clientID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var update dto.ClientUpdate

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	if update.Price != nil && (*update.Price < 0.01 || *update.Price > 999999) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Preço deve ser maior que zero"})
		return
	}

	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var clientUpdate models.Client

	if err := database.DB.First(&clientUpdate, uint(clientID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	if update.Name != nil {
		clientUpdate.Name = *update.Name
	}

	if update.Price != nil {
		clientUpdate.Price = *update.Price
	}

	if err := database.DB.Save(&clientUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto:" + err.Error()})
		return
	}

	response := dto.ClientResponse{
		ID:        clientUpdate.ID,
		Name:      clientUpdate.Name,
		Price:     clientUpdate.Price,
		CreatedAt: clientUpdate.CreatedAt,
		UpdatedAt: clientUpdate.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func GetClients(c *gin.Context) {

	var getClients []models.Client

	if err := database.DB.Find(&getClients).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produtos não encontrados"})
		return
	}

	var responses []dto.ClientResponse

	for _, p := range getClients {
		responses = append(responses, dto.ClientResponse{
			ID:        p.ID,
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	response := dto.ClientResponse{
		ID:        getClientId.ID,
		Name:      getClientId.Name,
		Price:     getClientId.Price,
		CreatedAt: getClientId.CreatedAt,
		UpdatedAt: getClientId.UpdatedAt,
	}

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
		c.JSON(http.StatusNotFound, gin.H{"error": "ID não encontrado"})
		return
	}

	if err := database.DB.Delete(&clientDelete, uint(ClientId)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto:" + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
