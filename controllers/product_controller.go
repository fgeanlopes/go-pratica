package controllers

import (
	"go-pratica/database"
	"go-pratica/dto"
	"go-pratica/models"
	"net/http"

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
	// id := c.Param("id")

	// ProductID, err := strconv.ParseUint(id, 10, 32)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
	// 	return
	// }

	// var product models.Product

}

func GetProducts(c *gin.Context) {

}

func GetProductByID(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
