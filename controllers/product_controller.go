package controllers

import (
	"go-pratica/database"
	"go-pratica/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar produto:" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)

}

func UpdateProduct(c *gin.Context) {

}

func GetProducts(c *gin.Context) {

}

func GetProductByID(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
