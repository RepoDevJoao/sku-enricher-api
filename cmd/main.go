package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/RepoDevJoao/sku-enricher-api/internal/handler"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	if os.Getenv("OPENAI_API_KEY") == "" {
		log.Fatal("OPENAI_API_KEY não encontrada")
	}

	r := gin.Default()

	r.GET("/health", handler.Health)
	r.POST("/enrich", handler.Enrich)
	r.POST("/normalize-sku", handler.NormalizeSKU)

	r.Run(":8080")
}