# Desafio 01 - Client-Server API

> Sistema client-server em Go para consulta e armazenamento de cotações de dólar em tempo real.

## Sobre

Sistema distribuído onde o **servidor** consome API externa de cotações e armazena no SQLite, e o **cliente** consulta o servidor e salva a cotação em arquivo local.

## Desafio

### Enunciado Original

> Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos, banco de dados e manipulação de arquivos com Go.
>
> Você precisará nos entregar dois sistemas em Go:
>
> - client.go
> - server.go

### Checklist de Requisitos

#### **Server.go**

- ✅ **Consumir API externa** - `https://economia.awesomeapi.com.br/json/last/USD-BRL`
- ✅ **Retornar JSON** - Formato com campo `bid`
- ✅ **Endpoint `/cotacao`** - Rota específica do desafio
- ✅ **Porta 8080** - Porta definida no enunciado
- ✅ **Context 200ms** - Timeout para chamada da API externa
- ✅ **Context 10ms** - Timeout para salvar no banco SQLite
- ✅ **Salvar no SQLite** - Registrar cada cotação recebida
- ✅ **Logs de erro** - Quando timeouts são excedidos

#### **Client.go**

- ✅ **Requisição HTTP** - Para o servidor solicitando cotação
- ✅ **Receber apenas `bid`** - Campo específico do JSON
- ✅ **Context 300ms** - Timeout para receber resultado do servidor
- ✅ **Salvar em arquivo** - `cotacao.txt` no formato "Dólar: {valor}"
- ✅ **Logs de erro** - Quando timeout é excedido

#### **Contextos e Timeouts**

- ✅ **3 contextos** - API externa (200ms), banco (10ms), cliente (300ms)
- ✅ **Logs de erro** - Para todos os timeouts excedidos
- ✅ **Tratamento adequado** - Erros não quebram o fluxo principal

## Configuração

### Pré-requisitos

- Go 1.21+
- SQLite3

### Instalação

```bash
go mod tidy
```

## Como Executar

### 1. Servidor

```bash
go run cmd/server/main.go
```

**Saída:**

```
Connecting to database...
Creating table...
Server initialized successfully
Server is running on port 8080
```

### 2. Cliente

```bash
go run cmd/client/main.go
```

**Saída:**

```
Quote saved -> Dólar: 5.1234
```

## Testando

### Health Check

```bash
curl http://localhost:8080/healthcheck
```

### Cotação

```bash
curl http://localhost:8080/cotacao
```

### Verificar Arquivo

```bash
cat cotacao.txt
# Dólar: 5.1234
```

### Verificar Banco

```bash
sqlite3 quotes.db
.tables
SELECT * FROM quotes;
.quit
```

## Timeouts

| Operação | Timeout | Comportamento |
|----------|---------|---------------|
| API Externa | 200ms | Log de erro se excedido |
| Banco SQLite | 10ms | Log de erro se excedido |
| Cliente HTTP | 300ms | Log de erro se excedido |

## Estrutura

```bash
fc-pos-golang-client-server-api/
├── cmd/
│   ├── client/main.go          # Cliente HTTP
│   └── server/main.go          # Servidor HTTP
├── internal/
│   ├── database/               # Camada de dados
│   ├── quote/                  # Domínio de cotação
│   └── server/                 # Servidor HTTP
├── .gitignore
├── go.mod
└── README.md
```

---

**Desenvolvido com ❤️ em Go**
