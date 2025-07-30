package controllers

import (
	"go-hospedes-api/config"
	"go-hospedes-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateHospede(context *gin.Context) {
	var hospede models.Hospede
	var err = context.ShouldBindJSON(&hospede)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var findErr = config.DB.First(&hospede, "phone = ?", hospede.Phone).Error
	if findErr == nil {
		context.JSON(http.StatusConflict, gin.H{"error": "Hospede com mesmo telefone já cadastrado"})
		return
	}

	config.DB.Create(&hospede)
	context.JSON(http.StatusCreated, hospede)
}

func GetAllHospedes(context *gin.Context) {
	var hospedes []models.Hospede
	config.DB.Find(&hospedes)
	context.JSON(http.StatusOK, hospedes)
}

func GetHospedesPaginated(context *gin.Context) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(context.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var hospedes []models.Hospede
	config.DB.Limit(limit).Offset(offset).Find(&hospedes)
	context.JSON(http.StatusOK, hospedes)
}

func GetHospedeById(context *gin.Context) {
	id := context.Param("id")
	var hospede models.Hospede
	result := config.DB.First(&hospede, id)

	if result.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Hospede não encontrado"})
		return
	}

	context.JSON(http.StatusOK, hospede)
}

func UpdateHospede(context *gin.Context) {
	id := context.Param("id")
	var hospede models.Hospede

	var findErr = config.DB.First(&hospede, id).Error
	if findErr != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Hospede não encontrado"})
		return
	}

	var input models.Hospede
	var bindErr = context.ShouldBindJSON(&input)
	if bindErr != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": bindErr.Error()})
	}

	config.DB.Model(&hospede).Updates(input)
	context.JSON(http.StatusOK, hospede)
}

func DeleteHsopede(context *gin.Context) {
	id := context.Param("id")
	var hospede models.Hospede

	var err = config.DB.First(&hospede, id).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Hospede não encontrado"})
		return
	}

	config.DB.Delete(&hospede)
	context.JSON(http.StatusOK, gin.H{"message": "Hóspede deletado com sucesso"})
}
