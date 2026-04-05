package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RepoDevJoao/sku-enricher-api/internal/service"
)

func Health(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func Enrich(c *gin.Context) {
	var input service.SKUInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	if input.ProductName == "" || input.Category == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_name e category são obrigatórios"})
		return
	}

	output, err := service.EnrichSKU(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func NormalizeSKU(c *gin.Context) {
	var input map[string]any

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	output, err := service.NormalizeSKU(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}