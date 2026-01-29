package main

import (
	"fmt"
	"go-pratica/database"
	"go-pratica/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":3000")
	fmt.Println("Hello, world!")
}
