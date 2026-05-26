# Notifications — Go SDK

> Notification management — view and manage notifications. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetAPIKeyNotification`

**`GET /api-key/notification`** — Get all notifications By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |
| `status` | query | `string` | No |  |

**Responses**

**200 OK** — `response.ListNotificationResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ListNotificationData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ListNotificationData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Notification]` |  |
| `total` | `integer` |  |

**`models.Notification`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `data` | `map[string]any` |  |
| `id` | `string` |  |
| `message` | `string` |  |
| `status` | `string` |  |
| `title` | `string` |  |
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
resp, err := client.Notifications.Notification.GetAPIKeyNotification(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyNotification`

**`DELETE /api-key/notification`** — Delete all notifications By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `status` | query | `string` | No |  |

**Responses**

**200 OK** — `response.SuccessResponse`

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` |  |
| `status` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Notifications.Notification.DeleteAPIKeyNotification(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyNotificationList`

**`PUT /api-key/notification/list`** — Mark list notification By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.UpdateListNotificationStatusByIdsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ids` | `array[string]` | Yes |  |
| `status` | `string` | Yes |  |

**Responses**

**200 OK** — `response.SuccessResponse`

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` |  |
| `status` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.UpdateListNotificationStatusByIdsRequest{
    Ids: "...",  // array[string]  // required
    Status: "...",  // string  // required
}
resp, err := client.Notifications.Notification.PutAPIKeyNotificationList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyNotificationList`

**`DELETE /api-key/notification/list`** — DeleteNotificationByIds By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.DeleteNotificationByIdsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ids` | `array[string]` | Yes |  |

**Responses**

**200 OK** — `response.SuccessResponse`

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` |  |
| `status` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.DeleteNotificationByIdsRequest{
    Ids: "...",  // array[string]  // required
}
resp, err := client.Notifications.Notification.DeleteAPIKeyNotificationList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyNotificationStatistics`

**`GET /api-key/notification/statistics`** — GetNotifyStatisticsByUserId By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.NotificationStatisticsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.NotificationStatisticsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.NotificationStatisticsData`**

| Field | Type | Description |
| --- | --- | --- |
| `total` | `integer` |  |
| `total_read` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Notifications.Notification.GetAPIKeyNotificationStatistics(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyNotificationID`

**`PUT /api-key/notification/{id}`** — Update notification by id By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Notification's id |
| `status` | query | `string` | Yes | Notification status |

**Request Body** — `request.UpdateNotificationByIdRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `status` | `string` | Yes |  |

**Responses**

**200 OK** — `response.SuccessResponse`

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` |  |
| `status` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.UpdateNotificationByIdRequest{
    Status: "...",  // string  // required
}
resp, err := client.Notifications.Notification.PutAPIKeyNotificationID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyNotificationID`

**`DELETE /api-key/notification/{id}`** — Delete notification by id By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Notification's id |

**Responses**

**200 OK** — `response.SuccessResponse`

| Field | Type | Description |
| --- | --- | --- |
| `message` | `string` |  |
| `status` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Notifications.Notification.DeleteAPIKeyNotificationID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
