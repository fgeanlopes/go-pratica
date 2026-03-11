package controllers

import (
	"go-pratica/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateVehicle(c *gin.Context) {
	var create dto.VehicleRequest

	if err := c.ShouldBindJSON(&create); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos:" + err.Error()})
		return
	}

	// vehicleCreate =: create.ToVehicleModel()

}
