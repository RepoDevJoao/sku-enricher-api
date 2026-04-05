# SKU Enricher API

Microservice backend in Go that enriches product/SKU data using AI — generating SEO titles, marketing descriptions, keywords, and ad copy aligned with e-commerce and digital advertising workflows.

Built as a demonstration of Go backend development with AI integration, directly inspired by real-world catalog enrichment pipelines.

---

## Stack

- **Go** + **Gin** — HTTP API
- **OpenAI API** (gpt-4o-mini) — AI content generation
- **godotenv** — environment configuration

---

## Endpoints

### `GET /health`
Returns API status.
```json
{ "status": "ok" }
```

---

### `POST /enrich`
Receives a product and returns AI-generated catalog content.

**Request:**
```json
{
  "product_name": "Tênis esportivo masculino",
  "category": "Calçados",
  "brand": "MoveFit",
  "features": ["solado antiderrapante", "tecido respirável"]
}
```

**Response:**
```json
{
  "seo_title": "Tênis Esportivo Masculino MoveFit - Conforto e Segurança",
  "marketing_description": "Ideal para quem busca desempenho...",
  "keywords": ["tênis masculino", "esportivo", "moveFit"],
  "short_ad_copy": "Conforto e leveza para todos os dias."
}
```

---

### `POST /normalize-sku`
Receives messy/inconsistent product data and returns a normalized, structured output.

**Request:**
```json
{
  "nome": "tenis msc movft",
  "cat": "calcados esportivos",
  "marca": "move fit"
}
```

**Response:**
```json
{
  "normalized_name": "Tênis Esportivo Masculino MoveFit",
  "suggested_category": "Calçados Esportivos",
  "tags": ["tênis", "masculino", "esportivo"]
}
```

---

## Running locally
```bash
# 1. Clone the repo
git clone https://github.com/RepoDevJoao/sku-enricher-api.git
cd sku-enricher-api

# 2. Set up environment
cp .env.example .env
# Add your OPENAI_API_KEY to .env

# 3. Install dependencies
go mod tidy

# 4. Run
go run cmd/main.go
```

---

## Project structure
```
sku-enricher/
├── cmd/
│   └── main.go          # entry point, routes
├── internal/
│   ├── handler/
│   │   └── sku.go       # HTTP handlers
│   └── service/
│       └── openai.go    # AI integration logic
├── .env                 # API keys (not committed)
└── go.mod
```

---

## Possible next steps

- Deploy on **Google Cloud Run**
- Store generations history in **Google Firestore**
- Add image prompt generation for product creatives
- Add SKU feed ingestion from CSV
- Implement **ADK (Agent Development Kit)** for agentic enrichment workflows
- Support **Agent2Agent (A2A) protocol** for multi-agent catalog pipelines