package routes

import (
	"go-pratica/controllers"

	"github.com/gin-gonic/gin"
)

func SetupClientsRoutes(c *gin.RouterGroup) {
	clients := c.Group("/clients")
	{
		clients.POST("/", controllers.CreateClient)
		clients.GET("/", controllers.GetClients)
		clients.GET("/:id", controllers.GetClientByID)
		clients.PUT("/:id", controllers.UpdateClient)
		clients.DELETE("/:id", controllers.DeleteClient)

		//TODO work on this
		// clients.GET("/:id/veiculos", controllers.GetVeiculosByCliente)
		// clients.GET("/:id/historico", controllers.GetHistoricoCliente)
	}
}
