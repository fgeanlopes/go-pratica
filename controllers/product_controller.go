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

func CreateProduct(c *gin.Context) {
	var create dto.ProductCreate

	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// Validação usando govalidator
	if _, err := govalidator.ValidateStruct(create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productCreate := models.Product{
		Name:  create.Name,
		Price: create.Price,
	}

	if err := database.DB.Create(&productCreate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto:" + err.Error()})
		return
	}

	response := dto.ProductResponse{
		ID:        productCreate.ID,
		Name:      productCreate.Name,
		Price:     productCreate.Price,
		CreatedAt: productCreate.CreatedAt,
		UpdatedAt: productCreate.UpdatedAt,
	}

	c.JSON(http.StatusCreated, response)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	productID, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var update dto.ProductUpdate

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// Validação usando govalidator
	if _, err := govalidator.ValidateStruct(update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var productUpdate models.Product

	if err := database.DB.First(&productUpdate, uint(productID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Produto não encontrado"})
		return
	}

	if update.Name != nil {
		productUpdate.Name = *update.Name
	}

	if update.Price != nil {
		productUpdate.Price = *update.Price
	}

	if err := database.DB.Save(&productUpdate).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar produto:" + err.Error()})
		return
	}

	response := dto.ProductResponse{
		ID:        productUpdate.ID,
		Name:      productUpdate.Name,
		Price:     productUpdate.Price,
		CreatedAt: productUpdate.CreatedAt,
		UpdatedAt: productUpdate.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)

	//Buscar o produto
	//Atualizar o model
	//Salvar o model no banco

	// var productUpdate models.Product

	// if update.Name != nil {
	// 	update.Name = *&update.Name
	// }

	// if update.Price != nil {
	// 	update.Price = *&update.Price
	// }

}

func GetProducts(c *gin.Context) {

}

func GetProductByID(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
