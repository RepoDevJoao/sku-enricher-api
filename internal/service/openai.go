package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SKUInput struct {
	ProductName string   `json:"product_name"`
	Category    string   `json:"category"`
	Brand       string   `json:"brand"`
	Features    []string `json:"features"`
}

type SKUOutput struct {
	SEOTitle    string   `json:"seo_title"`
	Description string   `json:"marketing_description"`
	Keywords    []string `json:"keywords"`
	AdCopy      string   `json:"short_ad_copy"`
}

func EnrichSKU(input SKUInput) (*SKUOutput, error) {
	prompt := fmt.Sprintf(`
Você é um especialista em catálogos de e-commerce e publicidade digital.
Dado o produto abaixo, gere conteúdo otimizado.

Produto: %s
Categoria: %s
Marca: %s
Características: %v

Responda APENAS com JSON válido, sem explicações, neste formato exato:
{
  "seo_title": "...",
  "marketing_description": "...",
  "keywords": ["...", "...", "..."],
  "short_ad_copy": "..."
}`, input.ProductName, input.Category, input.Brand, input.Features)

	body := map[string]any{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	var result map[string]any
	json.Unmarshal(respBytes, &result)

	choices := result["choices"].([]any)
	content := choices[0].(map[string]any)["message"].(map[string]any)["content"].(string)

	var output SKUOutput
	err = json.Unmarshal([]byte(content), &output)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear resposta da IA: %w", err)
	}

	return &output, nil
}

func NormalizeSKU(input map[string]any) (map[string]any, error) {
	raw, _ := json.Marshal(input)

	prompt := fmt.Sprintf(`
Você é um especialista em normalização de catálogos de e-commerce.
Dado este produto com dados bagunçados, normalize e estruture.

Dados brutos: %s

Responda APENAS com JSON válido, sem explicações, neste formato:
{
  "normalized_name": "...",
  "suggested_category": "...",
  "tags": ["...", "...", "..."]
}`, string(raw))

	body := map[string]any{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": prompt},
		},
	}

	bodyBytes, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro na requisição: %w", err)
	}
	defer resp.Body.Close()

	respBytes, _ := io.ReadAll(resp.Body)

	var result map[string]any
	json.Unmarshal(respBytes, &result)

	choices := result["choices"].([]any)
	content := choices[0].(map[string]any)["message"].(map[string]any)["content"].(string)

	var output map[string]any
	err = json.Unmarshal([]byte(content), &output)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear resposta: %w", err)
	}

	return output, nil
}