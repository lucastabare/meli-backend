# Guía rápida para levantar el backend (API estilo MercadoLibre).

Este documento es solo para el backend. Incluye requisitos, configuración, cómo ejecutar y cómo probar la API.

## Requisitos

Go 1.20+ (recomendado 1.21+)
Windows PowerShell (o bash/zsh en Linux/macOS)
(Opcional) curl/curl.exe para probar endpoints

## Estructura del proyecto (resumen)

```bash
backend/
  cmd/api/                 # main.go 
  internal/
    config/                # lectura de variables de entorno
    domain/                # modelos (Product, Seller, etc.)
    repository/jsonstore/  # Store + repos (persisten en archivos JSON)
    service/               # lógica de negocio
    transport/http/        # router chi + handlers + middlewares
  data/                    # *.json con datos de ejemplo
  go.mod / go.sum
```

## Variables de entorno

API_ADDR → dirección/puerto del server (default :8080)
DATA_DIR → carpeta donde están los JSON (default data)

### Windows (PowerShell):

```bash
$env:API_ADDR=":8080"
$env:DATA_DIR="data"
```

### Linux/macOS (bash/zsh):

```bash
export API_ADDR=":8080"
export DATA_DIR="data"
```

## Instalar dependencias 

```bash
go mod tidy
```
## Ejecutar en modo desarrollo

### Windows (PowerShell): 

```bash
# estando en backend/
$env:API_ADDR=":8080"
$env:DATA_DIR="data"
go run .\cmd\api
```

### Linux/macOS:

```bash
API_ADDR=":8080" DATA_DIR="data" go run ./cmd/api
```

El servidor quedará escuchando en http://localhost:8080.

## Prueba de la API

```bash
# Health
curl.exe http://localhost:8080/api/health

# Productos
curl.exe "http://localhost:8080/api/v1/products"
curl.exe "http://localhost:8080/api/v1/products?q=samsung&limit=10&offset=0"
curl.exe "http://localhost:8080/api/v1/products/MLA123"
curl.exe "http://localhost:8080/api/v1/products/MLA123/description"
curl.exe "http://localhost:8080/api/v1/products/MLA123/seller"

# Similares / Relacionados
curl.exe "http://localhost:8080/api/v1/products/MLA123/similar?limit=3"
curl.exe "http://localhost:8080/api/v1/products/MLA123/related?limit=3"

# Vendedor y métodos de pago
curl.exe "http://localhost:8080/api/v1/sellers/SELLER1"
curl.exe "http://localhost:8080/api/v1/payments/methods"

```
(En Linux/macOS, usá curl en lugar de curl.exe.)

## Datos de ejemplo 

La carpeta backend/data/ incluye:

products.json
sellers.json
payments.json

Podés editar estos JSON y la API reflejará los cambios sin reiniciar (se leen en cada request).

## Endpoints

GET /api/health → "ok"
GET /api/v1/products
Parámetros: q, limit, offset
Respuesta: { "total": number, "items": Product[] }

GET /api/v1/products/{id}
GET /api/v1/products/{id}/description
GET /api/v1/products/{id}/seller → { "seller_id": "..." }
GET /api/v1/products/{id}/similar?limit=K
GET /api/v1/products/{id}/related?limit=K
GET /api/v1/sellers/{id}
GET /api/v1/payments/methods

Errores:

404 cuando un recurso no existe → {"error":"not found"}
500 en errores internos → {"error":"internal error"}

# Correr con Docker

## Docker puro 
docker build -t meli-api .
docker run --rm -p 8080:8080 \
  -e ADDR=":8080" -e DATA_DIR="/app/data" \
  -v "$(pwd)/data:/app/data" \
  meli-api

## Docker compose 

docker compose up --build