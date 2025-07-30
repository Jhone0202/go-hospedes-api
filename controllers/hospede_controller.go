package controllers

import (
	"go-hospedes-api/config"
	"go-hospedes-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHospede(context *gin.Context){
	var hospede models.Hospede
	var err = context.ShouldBindJSON(&hospede);

	if err != nil {
		 context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		 return
	}
	
	config.DB.Create(&hospede)
	context.JSON(http.StatusCreated, hospede)
}

func GetAllHospedes(context *gin.Context){
	
}