# URL Shortener API

API para encurtar URLs construída com Go, Gin e GORM.

## Arquitetura

O projeto segue a arquitetura em camadas:

- **Models**: Definição das estruturas de dados
- **Repository** (`db/`): Acesso ao banco de dados
- **UseCase** (`route/usecase/`): Lógica de negócio
- **Handlers** (`route/handlers/`): Controllers HTTP
- **Routes** (`route/`): Definição das rotas

## Configuração

1. Configure as variáveis de ambiente no arquivo `.env`:

```env
APP_PORT=":8080"
DB_PATH="postgres://postgres:adminPassword@localhost:5432/url_shortener_db?sslmode=disable"
AUTO_MIGRATE="true"
```

2. Certifique-se de que o PostgreSQL está rodando.

## Executar

```bash
go run main.go
```

## Endpoints

### 1. Encurtar URL

```bash
POST /shorten
Content-Type: application/json

{
  "original_url": "https://www.exemplo.com.br/pagina-muito-longa"
}
```

**Resposta:**
```json
{
  "short_code": "abc123XY",
  "original_url": "https://www.exemplo.com.br/pagina-muito-longa",
  "short_url": "localhost:8080/abc123XY"
}
```

### 2. Redirecionar para URL Original

```bash
GET /:shortCode
```

Exemplo: `GET /abc123XY` redireciona para a URL original.

## Exemplo de Uso

```bash
# Encurtar uma URL
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"original_url": "https://www.google.com"}'

# Acessar URL encurtada (será redirecionado)
curl -L http://localhost:8080/abc123XY
```

## Docker Compose

Se houver `docker-compose.yml` configurado:

```bash
docker-compose up -d
```
