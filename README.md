# AIOZ AI Go Client

Go client SDK for the [AIOZ AI](https://aiozai.network) API. Auto-generated from the OpenAPI specification with full type safety.

## Features

- Full type safety for all 218+ API endpoints and 578+ models
- API Key configured once at client creation
- Automatic retry with exponential backoff for transient failures
- Configurable timeouts for standard and upload operations
- Domain-grouped service access (Models, Datasets, Competitions, etc.)
- Zero external HTTP dependencies beyond the Go standard library

## Installation

```bash
go get github.com/AIOZNetwork/aioz-ai-go-client@latest
```

Requires Go 1.22 or later.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"

    aiozai "github.com/AIOZNetwork/aioz-ai-go-client"
)

func main() {
    client, err := aiozai.NewClient(
        aiozai.WithAPIKey(os.Getenv("AIOZ_AI_API_KEY")),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    models := client.Models()
    _ = models // use models.Model, models.Training, etc.
    fmt.Println("client ready", ctx)
}
```

## Error Handling

```go
import "errors"

var apiErr *aiozai.AiozAPIError
if errors.As(err, &apiErr) {
    fmt.Printf("[%d] %s — %s %s\n", apiErr.StatusCode, apiErr.Message, apiErr.Method, apiErr.Endpoint)
}
```

## Configuration

```go
import "time"

client, err := aiozai.NewClient(
    aiozai.WithAPIKey(os.Getenv("AIOZ_AI_API_KEY")),
    aiozai.WithBaseURL("https://api.aiozai.network/api/v1"),
    aiozai.WithTimeout(60 * time.Second),
    aiozai.WithRetryConfig(&aiozai.RetryConfig{
        MaxRetries: 5,
        BaseDelay:  2 * time.Second,
        MaxDelay:   60 * time.Second,
    }),
)
```

## Service Groups

| Service | Access | Description | Reference |
| --- | --- | --- | --- |
| Models | `client.Models()` | AI model management | [docs/models.md](docs/models.md) |
| Datasets | `client.Datasets()` | Dataset management | [docs/datasets.md](docs/datasets.md) |
| Competitions | `client.Competitions()` | Competitions & submissions | [docs/competitions.md](docs/competitions.md) |
| Collections | `client.Collections()` | Curated collections | [docs/collections.md](docs/collections.md) |
| Discussions | `client.Discussions()` | Discussions & comments | [docs/discussions.md](docs/discussions.md) |
| Notifications | `client.Notifications()` | Notification system | [docs/notifications.md](docs/notifications.md) |
| Organizations | `client.Organizations()` | Organization management | [docs/organizations.md](docs/organizations.md) |
| Repositories | `client.Repositories()` | Repository operations | [docs/repositories.md](docs/repositories.md) |
| Storage | `client.Storage()` | Storage & uploads | [docs/storage.md](docs/storage.md) |
| Users | `client.Users()` | User management | [docs/users.md](docs/users.md) |
| Core | `client.Core()` | Core endpoints, search, offers | [docs/core.md](docs/core.md) |
| Public | `client.Public()` | Public endpoints (no auth) | [docs/public.md](docs/public.md) |

## License

Apache-2.0

<!-- SDK_GUIDE_START -->

---

## SDK Usage Guide

> Auto-generated from `swagger/sdk.json` — do not edit this section manually.
> Re-generate with `make guide` from the repo root.

### Authentication Setup

Initialize the client once with your API key and reuse it across all calls:

```go
import (
    "os"
    aiozai "github.com/AIOZNetwork/aioz-ai-go-client"
)

client, err := aiozai.NewClient(
    aiozai.WithAPIKey(os.Getenv("AIOZ_AI_API_KEY")),
)
if err != nil {
    log.Fatal(err)
}

// Use client.Models(), client.Datasets(), etc.
```

Obtain your API key from the [AIOZ AI dashboard](https://aiozai.network).

### Common Response Types

These types appear in error responses across all endpoints.

**`FailResponse`** (400 Bad Request)

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` | Human-readable error message |
| `errors` | `array[string]` | Field-level validation errors |

**`ErrorResponse`** (500 Internal Server Error)

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` | Internal error message |

<!-- SDK_GUIDE_END -->
