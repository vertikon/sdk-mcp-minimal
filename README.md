# SDK MCP Minimal

🚀 **SDK Público Minimalista para MCPs Vertikon**

## 📦 Features Essenciais

- ✅ **Logger**: Structured logging com Zap
- ✅ **HTTP**: Server com Chi v5 + CORS  
- ✅ **NATS**: Client básico para mensageria

## 🚀 Instalação

```go
go get github.com/vertikon/sdk-mcp-minimal
```

## 📋 Uso Rápido

### Logger
```go
import "github.com/vertikon/sdk-mcp-minimal/pkg/logger"

logger, _ := logger.NewProduction()
logger.Info("Hello MCP!")
```

### HTTP Server
```go
import (
	"github.com/vertikon/sdk-mcp-minimal/pkg/http"
)

router := http.NewRouter()
router.HandleFunc("/health", http.HealthHandler)

server := http.NewServer(":8080", router)
```

### NATS Client
```go
import "github.com/vertikon/sdk-mcp-minimal/pkg/nats"

client, _ := nats.NewClient("nats://localhost:4222")
defer client.Close()

client.Publish("mcp.events", map[string]interface{}{"id": "123"})
```

## 📊 Status

- ✅ Production Ready
- ✅ 5 Core Dependencies
- ✅ Public Repository
- ✅ Fast & Lightweight

## 🔗 Links

- GitHub: https://github.com/vertikon/sdk-mcp-minimal
- Issues: https://github.com/vertikon/sdk-mcp-minimal/issues

---

**Mantido pela equipe Vertikon** 🚀