package routes

import (
	"go-pratica/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
	router.PUT("/products/:id", controllers.UpdateProduct)
	router.DELETE("/products/:id", controllers.Delete)
}
