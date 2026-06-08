# Core — Go SDK

> Core platform operations — balance, search, offers, reactions, packages, and tasks. Requires: `x-api-key` header.

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
resp, err := client.Api-keys.Api-key.GetBalance(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutCommentsByID`

**`PUT /api-key/comments/{id}`** — Update comment

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Comment's id |

**Request Body** — `request.UpdateCommentRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | No |  |
| `tag_usernames` | `array[string]` | No |  |

**Responses**

**200 OK** — `response.CommentResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Comment` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Comment`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions` | `array[models.Reaction]` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `tag_usernames` | `array[string]` |  |
| `updated_at` | `string` |  |
| `violated` | `boolean` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.UpdateCommentRequest{
    Content: "...",  // string
    TagUsernames: "...",  // array[string]
}
resp, err := client.Commentss.Comments.PutCommentsByID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteCommentsByID`

**`DELETE /api-key/comments/{id}`** — Delete comment

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Comment's id |

**Responses**

**200 OK** — `response.CommentResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Comment` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Comment`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions` | `array[models.Reaction]` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `tag_usernames` | `array[string]` |  |
| `updated_at` | `string` |  |
| `violated` | `boolean` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Commentss.Comments.DeleteCommentsByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutDatasetByIDLike`

**`PUT /api-key/dataset/{id}/like`** — Like dataset

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Dataset's id |

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
resp, err := client.Reactions.Reaction.PutDatasetByIDLike(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDependencyLibRequest`

**`POST /api-key/dependency/lib-request`** — Create Dependency Package Request

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.CreateDependencyLibRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `package_name` | `string` | No |  |
| `reason` | `string` | No |  |
| `version` | `string` | No |  |

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
req := &models.CreateDependencyLibRequest{
    PackageName: "...",  // string
    Reason: "...",  // string
    Version: "...",  // string
}
resp, err := client.Dependencys.Dependency.PostDependencyLibRequest(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDependencyList`

**`POST /api-key/dependency/list`** — Get Dependency List

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.DependencyListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DependencyListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DependencyListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Dependency]` |  |
| `total` | `integer` |  |

**`models.Dependency`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `map[string]any` |  |
| `version` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Dependencys.Dependency.PostDependencyList(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetDependencyByID`

**`GET /api-key/dependency/{id}`** — Get Dependency

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Dependency's id |

**Responses**

**200 OK** — `response.DependencyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DependencyData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DependencyData`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |
| `version` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Dependencys.Dependency.GetDependencyByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetDiscussionByIDComments`

**`GET /api-key/discussion/{id}/comments`** — Get comment list

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Discussion's id |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.CommentListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CommentListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CommentListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Comment]` |  |
| `total` | `integer` |  |

**`models.Comment`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions` | `array[models.Reaction]` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `tag_usernames` | `array[string]` |  |
| `updated_at` | `string` |  |
| `violated` | `boolean` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Commentss.Comments.GetDiscussionByIDComments(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDiscussionByIDComments`

**`POST /api-key/discussion/{id}/comments`** — Create comment

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Discussion's id |

**Request Body** — `request.CreateCommentRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | No |  |
| `tag_usernames` | `array[string]` | No |  |

**Responses**

**200 OK** — `response.CommentResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Comment` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Comment`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions` | `array[models.Reaction]` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `tag_usernames` | `array[string]` |  |
| `updated_at` | `string` |  |
| `violated` | `boolean` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CreateCommentRequest{
    Content: "...",  // string
    TagUsernames: "...",  // array[string]
}
resp, err := client.Commentss.Comments.PostDiscussionByIDComments(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutItemsByIDReact`

**`PUT /api-key/items/{id}/react`** — React item

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Item's id |

**Request Body** — `request.ReactItemRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `itemName` | `string` | No |  |
| `reactName` | `string` | No |  |

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
req := &models.ReactItemRequest{
    Itemname: "...",  // string
    Reactname: "...",  // string
}
resp, err := client.Reactions.Reaction.PutItemsByIDReact(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutModelByIDLike`

**`PUT /api-key/model/{id}/like`** — Like model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

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
resp, err := client.Reactions.Reaction.PutModelByIDLike(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDTask`

**`POST /api-key/model/{id}/task`** — Distribute Task v2

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.DistributeTaskWithApiKeyRequest`

_No fields defined._

**Responses**

**200 OK** — `response.DistributeTaskResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `string` |  |
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
req := &models.DistributeTaskWithApiKeyRequest{
    
}
resp, err := client.Api-keys.Api-key.PostModelByIDTask(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostOfferInvite`

**`POST /api-key/offer/invite`** — Create invite to join organization

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.MakeInviteOfferRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `org_username` | `string` | Yes | OrgUsername is the username of the organization (required) |
| `role` | `string` | Yes | Role is the role of the new member (required) |
| `username` | `string` | Yes | Username is the username of the new member (required) |

**Responses**

**Error Responses**

| Status | Description |
| --- | --- |
| 201 | Created |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.MakeInviteOfferRequest{
    OrgUsername: "...",  // string  // required
    Role: "...",  // string  // required
    Username: "...",  // string  // required
}
resp, err := client.Offers.Offer.PostOfferInvite(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostOfferJoin`

**`POST /api-key/offer/join`** — Request to join organization

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.MakeJoinOfferRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `join_id` | `string` | Yes | OrgUsername is the username of the organization (required) |

**Responses**

**200 OK** — `response.MakeOfferResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.MakeOfferData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.MakeOfferData`**

| Field | Type | Description |
| --- | --- | --- |
| `offer` | `models.Offer` |  |

**`models.Offer`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `integer` |  |
| `created_by` | `string` |  |
| `exp_at` | `integer` |  |
| `id` | `string` |  |
| `org_username` | `string` |  |
| `role` | `string` |  |
| `type` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.MakeJoinOfferRequest{
    JoinID: "...",  // string  // required
}
resp, err := client.Offers.Offer.PostOfferJoin(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutOfferByOfferIDAccept`

**`PUT /api-key/offer/{offer_id}/accept`** — Accept offer

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `offer_id` | path | `string` | Yes | offer_id |

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
resp, err := client.Offers.Offer.PutOfferByOfferIDAccept(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutOfferByOfferIDDeny`

**`PUT /api-key/offer/{offer_id}/deny`** — Deny offer

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `offer_id` | path | `string` | Yes | offer_id |

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
resp, err := client.Offers.Offer.PutOfferByOfferIDDeny(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetOfferByOfferIDResend`

**`GET /api-key/offer/{offer_id}/resend`** — Resend offer

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `offer_id` | path | `string` | Yes | offer_id |

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
resp, err := client.Offers.Offer.GetOfferByOfferIDResend(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteOfferByOfferIDRevoke`

**`DELETE /api-key/offer/{offer_id}/revoke`** — Revoke offer

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `offer_id` | path | `string` | Yes | offer_id |

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
resp, err := client.Offers.Offer.DeleteOfferByOfferIDRevoke(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPackageApiBuy`

**`POST /api-key/package/api/buy`** — Buy Api Key Package

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.BuyApiKeyPackageRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_key_id` | `string` | No |  |
| `model_id` | `string` | No |  |
| `quantity_uses` | `integer` | No |  |

**Responses**

**200 OK** — `response.ApiPackageResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ApiPackageData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ApiPackageData`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `discount` | `number` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `quantity_uses` | `integer` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.BuyApiKeyPackageRequest{
    APIKeyID: "...",  // string
    ModelID: "...",  // string
    QuantityUses: "...",  // integer
}
resp, err := client.Packages.Package.PostPackageApiBuy(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPackageBuyCost`

**`POST /api-key/package/buy/cost`** — Calculate Cost To Buy Package

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.BuyPlaygroundPackageRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `model_id` | `string` | No |  |
| `quantity_uses` | `integer` | No |  |

**Responses**

**200 OK** — `response.EstimateCostResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.EstimateCostData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.EstimateCostData`**

| Field | Type | Description |
| --- | --- | --- |
| `cost` | `number` |  |
| `symbol` | `string` |  |
| `unit` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.BuyPlaygroundPackageRequest{
    ModelID: "...",  // string
    QuantityUses: "...",  // integer
}
resp, err := client.Packages.Package.PostPackageBuyCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPackageMasterBuyCost`

**`POST /api-key/package/master/buy/cost`** — Calculate Cost To Buy Master Package

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.BuyPlaygroundMasterPackageRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `quantity_uses` | `integer` | No |  |

**Responses**

**200 OK** — `response.EstimateCostResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.EstimateCostData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.EstimateCostData`**

| Field | Type | Description |
| --- | --- | --- |
| `cost` | `number` |  |
| `symbol` | `string` |  |
| `unit` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.BuyPlaygroundMasterPackageRequest{
    QuantityUses: "...",  // integer
}
resp, err := client.Packages.Package.PostPackageMasterBuyCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPackagePlaygroundBuy`

**`POST /api-key/package/playground/buy`** — Unlock Playground Package

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.BuyPlaygroundPackageRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `model_id` | `string` | No |  |
| `quantity_uses` | `integer` | No |  |

**Responses**

**200 OK** — `response.ApiPackageResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ApiPackageData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ApiPackageData`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `discount` | `number` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `quantity_uses` | `integer` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.BuyPlaygroundPackageRequest{
    ModelID: "...",  // string
    QuantityUses: "...",  // integer
}
resp, err := client.Packages.Package.PostPackagePlaygroundBuy(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPackageByID`

**`GET /api-key/package/{id}`** — Get Api Package

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | ApiPackage's id |

**Responses**

**200 OK** — `response.ApiPackageResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ApiPackageData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ApiPackageData`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `discount` | `number` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `quantity_uses` | `integer` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Packages.Package.GetPackageByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPermission`

**`GET /api-key/permission`** — Get Api Key Permission

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.GetApiKeyPermissionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetApiKeyPermissionData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetApiKeyPermissionData`**

| Field | Type | Description |
| --- | --- | --- |
| `api_key_packages` | `array[models.ApiKeyPackage]` |  |
| `limit_models` | `boolean` |  |
| `models` | `array[string]` |  |

**`models.ApiKeyPackage`**

| Field | Type | Description |
| --- | --- | --- |
| `model_id` | `string` |  |
| `uses_remaining` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Api-keys.Api-key.GetPermission(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPlatformTaskDatasetByIDTraining`

**`GET /api-key/platform/task/dataset/{id}/training`** — Get List Platform Task By DatasetId And UserId

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Dataset's Id |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.QueueTaskListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.QueueTaskListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.QueueTaskListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.QueueTask]` |  |
| `total` | `integer` |  |

**`models.QueueTask`**

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `object_id` | `string` |  |
| `object_name` | `string` |  |
| `result` | `map[string]any` |  |
| `state` | `string` |  |
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
resp, err := client.Platforms.Task.GetPlatformTaskDatasetByIDTraining(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPlatformTaskModelByIDVerify`

**`GET /api-key/platform/task/model/{id}/verify`** — Get List Platform Task By ModelId And UserId

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's Id |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.QueueTaskListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.QueueTaskListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.QueueTaskListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.QueueTask]` |  |
| `total` | `integer` |  |

**`models.QueueTask`**

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `object_id` | `string` |  |
| `object_name` | `string` |  |
| `result` | `map[string]any` |  |
| `state` | `string` |  |
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
resp, err := client.Platforms.Task.GetPlatformTaskModelByIDVerify(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPlatformTaskByID`

**`GET /api-key/platform/task/{id}`** — Get Platform Task By Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Task's Id |

**Responses**

**200 OK** — `response.QueueTaskResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.QueueTask` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.QueueTask`**

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `object_id` | `string` |  |
| `object_name` | `string` |  |
| `result` | `map[string]any` |  |
| `state` | `string` |  |
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
resp, err := client.Platforms.Task.GetPlatformTaskByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetSearch`

**`GET /api-key/search`** — Multiple Search

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `search` | query | `string` | Yes |  |

**Responses**

**200 OK** — `response.SearchResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.SearchResponseData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.SearchResponseData`**

| Field | Type | Description |
| --- | --- | --- |
| `competition` | `response.CompetitionListData` |  |
| `dataset` | `response.DatasetListData` |  |
| `model` | `response.ModelListData` |  |
| `organization` | `response.OrganizationListData` |  |
| `user` | `response.UserListData` |  |

**`response.CompetitionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Competition]` |  |
| `total` | `integer` |  |

**`models.Competition`**

| Field | Type | Description |
| --- | --- | --- |
| `author_id` | `string` |  |
| `category` | `string` |  |
| `code` | `string` |  |
| `cover` | `string` |  |
| `created_at` | `string` |  |
| `data` | `string` |  |
| `description` | `string` |  |
| `end_date` | `string` |  |
| `final_result_mode` | `string` |  |
| `id` | `string` |  |
| `joined` | `boolean` |  |
| `launched` | `boolean` |  |
| `max_daily_private_submissions` | `integer` |  |
| `overview` | `string` |  |
| `owner` | `models.Owner` |  |
| `participants` | `integer` |  |
| `path` | `string` |  |
| `permission` | `map[string]any` |  |
| `private_leaderboard_release_date` | `string` |  |
| `private_submissions_remaining` | `integer` |  |
| `prize_distribution_method` | `string` |  |
| `registration_deadline` | `string` |  |
| `reward_type` | `string` |  |
| `rules` | `string` |  |
| `start_date` | `string` |  |
| `submission_deadline` | `string` |  |
| `submissions` | `integer` |  |
| `tags` | `array[string]` |  |
| `thumbnail` | `string` |  |
| `time_zone_config` | `map[string]any` |  |
| `title` | `string` |  |
| `total_prize_pool` | `number` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `visibility` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`response.DatasetListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Dataset]` |  |
| `total` | `integer` |  |

**`models.Dataset`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `author_id` | `string` |  |
| `cover` | `string` |  |
| `create_by` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `discussions_count` | `integer` |  |
| `downloads_count` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `is_official` | `boolean` |  |
| `is_released` | `boolean` |  |
| `is_verified` | `boolean` |  |
| `likes_count` | `integer` |  |
| `metadata` | `models.DatasetMetadata` |  |
| `name` | `string` |  |
| `price` | `number` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `username` | `string` |  |
| `visibility` | `string` |  |

**`models.DatasetMetadata`**

| Field | Type | Description |
| --- | --- | --- |
| `dataset_id` | `string` |  |
| `id` | `string` |  |
| `language` | `array[string]` | en, vi |
| `license` | `string` |  |
| `pretty_name` | `string` |  |
| `size_category` | `string` |  |
| `tags` | `array[string]` | art |
| `task_categories` | `array[string]` | text-to-image |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**`response.ModelListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Model]` |  |
| `total` | `integer` |  |

**`models.Model`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `author_id` | `string` |  |
| `commit_hash` | `string` |  |
| `cover` | `string` |  |
| `create_by` | `string` |  |
| `created_at` | `string` |  |
| `dependency_id` | `string` |  |
| `description` | `string` |  |
| `discussions_count` | `integer` |  |
| `downloads_count` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `is_official` | `boolean` |  |
| `is_released` | `boolean` |  |
| `is_verified` | `boolean` |  |
| `likes_count` | `integer` |  |
| `model_metadata` | `models.ModelMetadata` |  |
| `name` | `string` |  |
| `playground_count` | `integer` |  |
| `price` | `number` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `task_reviews_count` | `integer` |  |
| `task_reviews_point` | `number` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `username` | `string` |  |
| `visibility` | `string` |  |

**`models.ModelMetadata`**

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `license` | `string` |  |
| `model_id` | `string` |  |
| `pretty_name` | `string` |  |
| `task` | `string` |  |

**`models.Reaction`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `name` | `string` |  |
| `owner` | `models.Owner` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.ReactionStats`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `name` | `string` |  |

**`response.OrganizationListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.OrganizationInfo]` |  |
| `total` | `integer` |  |

**`models.OrganizationInfo`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `full_name` | `string` |  |
| `id` | `string` |  |
| `org_type` | `string` |  |
| `permission` | `map[string]any` |  |
| `type` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |

**`response.UserListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.UserInfo]` |  |
| `total` | `integer` |  |

**`models.UserInfo`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `bio` | `string` |  |
| `fullname` | `string` |  |
| `github_link` | `string` |  |
| `home_page` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `username` | `string` |  |
| `visibility` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Searchs.Search.GetSearch(ctx)
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
| `total_cost` | `integer` |  |
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
req := &models.GetApiKeyStatisticsByModelIdRequest{
    From: "...",  // string
    To: "...",  // string
}
resp, err := client.Api-keys.Api-key.PostStatistics(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostTask`

**`POST /api-key/task`** — Distribute Task

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.DistributeTaskRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `files` | `array[object]` | No |  |
| `input_params` | `map[string]any` | No |  |
| `model_id` | `string` | No |  |

**Responses**

**200 OK** — `response.DistributeTaskResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `string` |  |
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
req := &models.DistributeTaskRequest{
    Files: "...",  // array[object]
    InputParams: "...",  // map[string]any
    ModelID: "...",  // string
}
resp, err := client.Api-keys.Api-key.PostTask(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetTaskHistories`

**`GET /api-key/task/histories`** — Get Tasks Histories

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.ApiKeyHistoryListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ApiKeyHistoryListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ApiKeyHistoryListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.ApiKeyHistories]` |  |
| `total` | `integer` |  |

**`models.ApiKeyHistories`**

| Field | Type | Description |
| --- | --- | --- |
| `cost` | `number` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `input_data` | `map[string]any` |  |
| `input_format` | `map[string]any` |  |
| `message` | `string` |  |
| `output_format` | `map[string]any` |  |
| `result` | `models.InferenceResult` |  |
| `status` | `string` |  |
| `task_id` | `string` |  |
| `updated_at` | `string` |  |

**`models.InferenceResult`**

| Field | Type | Description |
| --- | --- | --- |
| `error` | `map[string]any` |  |
| `result` | `map[string]any` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Api-keys.Api-key.GetTaskHistories(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteTaskByIDCancel`

**`DELETE /api-key/task/{id}/cancel`** — Cancel Task

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Task's id |

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
resp, err := client.Api-keys.Api-key.DeleteTaskByIDCancel(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetTaskByIDDetail`

**`GET /api-key/task/{id}/detail`** — Get Api Key Log By TaskId

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Task's id |

**Responses**

**200 OK** — `response.ApiKeyHistoryResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ApiKeyHistories` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ApiKeyHistories`**

| Field | Type | Description |
| --- | --- | --- |
| `cost` | `number` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `input_data` | `map[string]any` |  |
| `input_format` | `map[string]any` |  |
| `message` | `string` |  |
| `output_format` | `map[string]any` |  |
| `result` | `models.InferenceResult` |  |
| `status` | `string` |  |
| `task_id` | `string` |  |
| `updated_at` | `string` |  |

**`models.InferenceResult`**

| Field | Type | Description |
| --- | --- | --- |
| `error` | `map[string]any` |  |
| `result` | `map[string]any` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Api-keys.Api-key.GetTaskByIDDetail(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetTrainingTaskByID`

**`GET /api-key/training-task/{id}`** — Get Training Task

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Training Task's id |

**Responses**

**200 OK** — `response.TrainingTaskResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.TrainingTaskData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.TrainingTaskData`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `metadata` | `string` |  |
| `name` | `string` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Trainings.Task.GetTrainingTaskByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetTrainingTaskList`

**`GET /api-key/training/task/list`** — Get Training Task List

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.TrainingTaskListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.TrainingTaskListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.TrainingTaskListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.TrainingTask]` |  |
| `total` | `integer` |  |

**`models.TrainingTask`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `metadata` | `map[string]any` |  |
| `name` | `string` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Trainings.Task.GetTrainingTaskList(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
