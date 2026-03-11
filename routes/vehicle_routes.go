package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupVehicleRoutes(c *gin.RouterGroup) {
	vehicles := c.Group("/vehicles")
	{
		vehicles.POST("/", controllers.CreateVehicle)
		// vehicles.GET("/", controllers.GetVehicles)
		// vehicles.GET("/:id", controllers.GetVehicleByID)
		// vehicles.GET("/plate/:plate", controllers.GetVehicleByPlate)
		// vehicles.PUT("/:id", controllers.UpdateVehicle)
		// vehicles.DELETE("/:id", controllers.DeleteVehicle)
	}
}
