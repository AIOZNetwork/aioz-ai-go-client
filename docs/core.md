# Core — Go SDK

> Core platform operations — balance, info, and statistics. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetBalance`

**`GET /api-key/balance`** — Get Api Key Balance

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.WalletWithAddressResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.WalletWithAddress` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.WalletWithAddress`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `string` |  |
| `debt` | `string` |  |
| `earnings` | `string` |  |
| `free_balance` | `string` |  |
| `wallet_address` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Core().Core.GetBalance(core.NewGetBalanceParams().WithContext(ctx), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetInfo`

**`GET /api-key/info`** — Get Api Key Info

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.LiteApiKeyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteApiKey` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteApiKey`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `api_key` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `model_id` | `string` |  |
| `name` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Core().Core.GetInfo(core.NewGetInfoParams().WithContext(ctx), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostStatistics`

**`POST /api-key/statistics`** — Get Api Key Statistics

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.GetApiKeyStatisticsByModelIdRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `to` | `string` | No | 2023-05-07 15:04:05 |

**Responses**

**200 OK** — `response.GetTaskStatisticsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetTaskStatisticsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetTaskStatisticsData`**

| Field | Type | Description |
| --- | --- | --- |
| `total_cost` | `number` |  |
| `total_failed` | `integer` |  |
| `total_request` | `integer` |  |
| `total_success` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestGetAPIKeyStatisticsByModelIDRequest{
    From: "...",  // string
    To: "...",  // string
}
resp, err := client.Core().Core.PostStatistics(core.NewPostStatisticsParams().WithContext(ctx).WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
