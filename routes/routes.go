package routes

import (
	"go-hospedes-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/hospedes")
	{
		api.POST("/", controllers.CreateHospede)
		api.GET("/", controllers.GetAllHospedes)
		api.GET("/paginated", controllers.GetHospedesPaginated)
		api.GET("/:id", controllers.GetHospedeById)
		api.PUT("/:id", controllers.UpdateHospede)
		api.DELETE("/:id", controllers.DeleteHsopede)
	}
}
