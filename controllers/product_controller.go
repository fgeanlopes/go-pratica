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
}

func GetProducts(c *gin.Context) {

	var getProducts []models.Product

	if err := database.DB.Find(&getProducts).Error; err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": "Produtos não encontrados"})
		return
	}

	var responses []dto.ProductResponse

	for _, p := range getProducts {
		responses = append(responses, dto.ProductResponse{
			ID:        p.ID,
			Name:      p.Name,
			Price:     p.Price,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, responses)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	productId, err := strconv.ParseInt(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var getProductId models.Product

	if err := database.DB.First(&getProductId, uint(productId)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
	}

	response := dto.ProductResponse{
		ID:        getProductId.ID,
		Name:      getProductId.Name,
		Price:     getProductId.Price,
		CreatedAt: getProductId.CreatedAt,
		UpdatedAt: getProductId.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}

func Delete(c *gin.Context) {
	id := c.Param("id")

	ProductId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var productDelete models.Product

	// TODO otimizar para evitar 2 round-trips
	if err := database.DB.First(&productDelete, uint(ProductId)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID não encontrado"})
		return
	}

	if err := database.DB.Delete(&productDelete, uint(ProductId)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar produto:" + err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
