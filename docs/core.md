# Core — Go SDK

> Core platform operations — balance, search, offers, reactions, packages, and tasks. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetAPIKeyBalance`

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
| `balance` | `number` |  |
| `debt` | `number` |  |
| `earnings` | `number` |  |
| `free_balance` | `number` |  |
| `wallet_address` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Api-keys.Api-key.GetAPIKeyBalance(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyCommentsID`

**`PUT /api-key/comments/{id}`** — Update comment By Api Key

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
resp, err := client.Commentss.Comments.PutAPIKeyCommentsID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyCommentsID`

**`DELETE /api-key/comments/{id}`** — Delete comment By Api Key

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
resp, err := client.Commentss.Comments.DeleteAPIKeyCommentsID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyDatasetIDLike`

**`PUT /api-key/dataset/{id}/like`** — Like dataset By Api Key

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
resp, err := client.Reactions.Reaction.PutAPIKeyDatasetIDLike(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyDependencyLibRequest`

**`POST /api-key/dependency/lib-request`** — Create Dependency Package Request By Api Key

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
resp, err := client.Dependencys.Dependency.PostAPIKeyDependencyLibRequest(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyDependencyList`

**`POST /api-key/dependency/list`** — Get Dependency List By Api Key

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
resp, err := client.Dependencys.Dependency.PostAPIKeyDependencyList(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyDependencyID`

**`GET /api-key/dependency/{id}`** — Get Dependency By Api Key

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
resp, err := client.Dependencys.Dependency.GetAPIKeyDependencyID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyDiscussionIDComments`

**`GET /api-key/discussion/{id}/comments`** — Get comment list By Api Key

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
resp, err := client.Commentss.Comments.GetAPIKeyDiscussionIDComments(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyDiscussionIDComments`

**`POST /api-key/discussion/{id}/comments`** — Create comment By Api Key

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
resp, err := client.Commentss.Comments.PostAPIKeyDiscussionIDComments(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyItemsIDReact`

**`PUT /api-key/items/{id}/react`** — React item By Api Key

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
resp, err := client.Reactions.Reaction.PutAPIKeyItemsIDReact(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyModelIDLike`

**`PUT /api-key/model/{id}/like`** — Like model By Api Key

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
resp, err := client.Reactions.Reaction.PutAPIKeyModelIDLike(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyModelIDTask`

**`POST /api-key/model/{id}/task`** — Distribute Task v2 (Api-Key)

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
resp, err := client.Api-keys.Api-key.PostAPIKeyModelIDTask(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOfferInvite`

**`POST /api-key/offer/invite`** — Create invite to join organization By Api Key

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
resp, err := client.Offers.Offer.PostAPIKeyOfferInvite(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOfferJoin`

**`POST /api-key/offer/join`** — Request to join organization By Api Key

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
resp, err := client.Offers.Offer.PostAPIKeyOfferJoin(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyOfferOffer_idAccept`

**`PUT /api-key/offer/{offer_id}/accept`** — Accept offer By Api Key

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
resp, err := client.Offers.Offer.PutAPIKeyOfferOffer_idAccept(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutAPIKeyOfferOffer_idDeny`

**`PUT /api-key/offer/{offer_id}/deny`** — Deny offer By Api Key

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
resp, err := client.Offers.Offer.PutAPIKeyOfferOffer_idDeny(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOfferOffer_idResend`

**`GET /api-key/offer/{offer_id}/resend`** — Resend offer By Api Key

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
resp, err := client.Offers.Offer.GetAPIKeyOfferOffer_idResend(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyOfferOffer_idRevoke`

**`DELETE /api-key/offer/{offer_id}/revoke`** — Revoke offer By Api Key

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
resp, err := client.Offers.Offer.DeleteAPIKeyOfferOffer_idRevoke(ctx, "<offer_id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyPackageAPIBuy`

**`POST /api-key/package/api/buy`** — Buy Api Key Package By Api Key

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
resp, err := client.Packages.Package.PostAPIKeyPackageAPIBuy(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyPackageBuyCost`

**`POST /api-key/package/buy/cost`** — Calculate Cost To Buy Package By Api Key

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
resp, err := client.Packages.Package.PostAPIKeyPackageBuyCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyPackageMasterBuyCost`

**`POST /api-key/package/master/buy/cost`** — Calculate Cost To Buy Master Package By Api Key

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
resp, err := client.Packages.Package.PostAPIKeyPackageMasterBuyCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyPackagePlaygroundBuy`

**`POST /api-key/package/playground/buy`** — Unlock Playground Package By Api Key

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
resp, err := client.Packages.Package.PostAPIKeyPackagePlaygroundBuy(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyPackageID`

**`GET /api-key/package/{id}`** — Get Api Package By Api Key

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
resp, err := client.Packages.Package.GetAPIKeyPackageID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyPermission`

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
resp, err := client.Api-keys.Api-key.GetAPIKeyPermission(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyPlatformTaskDatasetIDTraining`

**`GET /api-key/platform/task/dataset/{id}/training`** — Get List Platform Task By DatasetId And UserId By Api Key

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
resp, err := client.Platforms.Task.GetAPIKeyPlatformTaskDatasetIDTraining(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyPlatformTaskModelIDVerify`

**`GET /api-key/platform/task/model/{id}/verify`** — Get List Platform Task By ModelId And UserId By Api Key

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
resp, err := client.Platforms.Task.GetAPIKeyPlatformTaskModelIDVerify(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyPlatformTaskID`

**`GET /api-key/platform/task/{id}`** — Get Platform Task By Id By Api Key

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
resp, err := client.Platforms.Task.GetAPIKeyPlatformTaskID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeySearch`

**`GET /api-key/search`** — Multiple Search By Api Key

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
| `competition` | `array[object]` |  |
| `dataset` | `array[object]` |  |
| `model` | `array[object]` |  |
| `organization` | `array[object]` |  |
| `user` | `array[object]` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Searchs.Search.GetAPIKeySearch(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyStatistics`

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
resp, err := client.Api-keys.Api-key.PostAPIKeyStatistics(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyTask`

**`POST /api-key/task`** — Distribute Task (Api-Key)

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
resp, err := client.Api-keys.Api-key.PostAPIKeyTask(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyTaskHistories`

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
resp, err := client.Api-keys.Api-key.GetAPIKeyTaskHistories(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyTaskIDCancel`

**`DELETE /api-key/task/{id}/cancel`** — Cancel Task By Api Key

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
resp, err := client.Api-keys.Api-key.DeleteAPIKeyTaskIDCancel(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyTaskIDDetail`

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
resp, err := client.Api-keys.Api-key.GetAPIKeyTaskIDDetail(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyTrainingTaskID`

**`GET /api-key/training-task/{id}`** — Get Training Task By Api Key

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
resp, err := client.Trainings.Task.GetAPIKeyTrainingTaskID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyTrainingTaskList`

**`GET /api-key/training/task/list`** — Get Training Task List By Api Key

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
resp, err := client.Trainings.Task.GetAPIKeyTrainingTaskList(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
