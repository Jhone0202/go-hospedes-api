package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Wellcome(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", nil)
}
