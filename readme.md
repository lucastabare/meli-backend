# Backend – MercadoLibre Product API (Go)

API para una PDP (product detail page) estilo MercadoLibre.
Persistencia en JSON local y arquitectura en capas.

> **Autor:** Desarrollado por **Lucas Tabaré**.

## Cómo ejecutar: ver backend/run.md.

## ✨ Features
REST API versionada en /api/v1

Productos: listado + búsqueda, detalle, descripción, similares y relacionados

Vendedores: detalle

Pagos: métodos disponibles

CORS habilitado para desarrollo

## 🧰 Stack
Go 1.20+

chi v5 (router)

Datos en /data/*.json

## 🗂️ Estructura

```bash
backend/
  cmd/api/                 # main.go
  internal/
    config/                # env vars (API_ADDR, DATA_DIR)
    domain/                # modelos
    repository/jsonstore/  # store + repos (JSON)
    service/               # lógica negocio
    transport/http/        # router + handlers
  data/                    # products.json, sellers.json, payments.json
  run.md
  go.mod / go.sum

```

## ⚙️ Config

API_ADDR (default :8080)
DATA_DIR (default data)

## 🔌 Endpoints

GET /api/health → "ok"
GET /api/v1/products?q=&limit=&offset= → { total, items[] }
GET /api/v1/products/{id}
GET /api/v1/products/{id}/description
GET /api/v1/products/{id}/seller → { seller_id }
GET /api/v1/products/{id}/similar?limit=K
GET /api/v1/products/{id}/related?limit=K
GET /api/v1/sellers/{id}
GET /api/v1/payments/methods