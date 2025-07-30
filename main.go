package main

import (
	"go-hospedes-api/config"
	"go-hospedes-api/models"
	"go-hospedes-api/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	config.Connect()
	models.Migate(config.DB)

	engine := gin.Default()
	routes.SetupRoutes(engine)

	engine.Run(":8080")
}