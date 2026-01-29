package controllers

import (
	"go-pratica/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

}

func UpdateProduct(c *gin.Context) {

}

func GetProducts(c *gin.Context) {

}

func GetProductByID(c *gin.Context) {

}

func Delete(c *gin.Context) {

}
