package routes

import (
	"go-hospedes-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine){
	api := r.Group("/api/hospedes")
	{
		api.POST("/", controllers.CreateHospede)
	}
}