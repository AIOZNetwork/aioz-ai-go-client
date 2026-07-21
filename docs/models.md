# Models — Go SDK

> Model management — create, train, deploy, and version models. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `PostModel`

**`POST /api-key/model`** — Create Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.CreateModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `author_id` | `string` | No |  |
| `cover` | `string` | Yes |  |
| `dependency_id` | `string` | Yes |  |
| `description` | `string` | Yes |  |
| `language` | `array[string]` | No | en, vi |
| `library` | `array[string]` | No | tf |
| `license` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `pretty_name` | `string` | No |  |
| `price` | `number` | No |  |
| `task` | `string` | No |  |
| `thumbnail` | `string` | Yes |  |
| `visibility` | `string` | Yes | One of: `public`, `private` |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Model` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestCreateModelRequest{
    AuthorID: "...",  // string
    Cover: "...",  // string  // required
    DependencyID: "...",  // string  // required
    Description: "...",  // string  // required
    Language: "...",  // array[string]
}
resp, err := client.Models().Model.PostModel(model.NewPostModelParams().WithContext(ctx).WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelList`

**`POST /api-key/model/list`** — Get Model List

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.GetModelListRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `filter_by` | `string` | No | One of: `public`, `community`, `author`, `official`, `permission`, `playground` |
| `language` | `array[string]` | No |  |
| `library` | `array[string]` | No |  |
| `license` | `string` | No |  |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `asc`, `desc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `likes`, `downloads`, `created`, `created_oldest`, `modified`, `name`, `size` |
| `tag_type` | `string` | No |  |
| `task` | `string` | No |  |

**Responses**

**200 OK** — `response.ModelListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestGetModelListRequest{
    FilterBy: "...",  // string
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
}
resp, err := client.Models().Model.PostModelList(model.NewPostModelListParams().WithContext(ctx).WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelListByAuthorByUsername`

**`POST /api-key/model/list-by-author/{username}`** — Get Model List By User

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | User's username |

**Request Body** — `request.GetModelListByAuthorRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `language` | `array[string]` | No |  |
| `library` | `array[string]` | No |  |
| `license` | `string` | No |  |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `asc`, `desc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `likes`, `downloads`, `created`, `created_oldest`, `modified`, `name`, `size` |
| `tag_type` | `string` | No |  |
| `task` | `string` | No |  |

**Responses**

**200 OK** — `response.ModelListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestGetModelListByAuthorRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Models().Model.PostModelListByAuthorByUsername(model.NewPostModelListByAuthorByUsernameParams().WithContext(ctx).WithUsername("<username>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelMatchingTags`

**`POST /api-key/model/matching-tags`** — MatchingModels Tags

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.MatchingModelsTagsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `language` | `array[string]` | No |  |
| `library` | `array[string]` | No |  |
| `license` | `string` | No |  |
| `tag_type` | `string` | No |  |
| `task` | `string` | No |  |

**Responses**

**200 OK** — `response.MatchingModelsTagsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.MatchingModelsTagsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.MatchingModelsTagsData`**

| Field | Type | Description |
| --- | --- | --- |
| `languages` | `array[string]` |  |
| `libraries` | `array[string]` |  |
| `licenses` | `array[string]` |  |
| `tasks` | `array[string]` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestMatchingModelsTagsRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    TagType: "...",  // string
    Task: "...",  // string
}
resp, err := client.Models().Model.PostModelMatchingTags(model.NewPostModelMatchingTagsParams().WithContext(ctx).WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelOrganizationByOrg`

**`GET /api-key/model/organization/{org}`** — Get List Model By Org Username

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | Org's username |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.ModelListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelOrganizationByOrg(model.NewGetModelOrganizationByOrgParams().WithContext(ctx).WithOrg("<org>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByID`

**`GET /api-key/model/{id}`** — Get Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Model` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelByID(model.NewGetModelByIDParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutModelByID`

**`PUT /api-key/model/{id}`** — Update Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.UpdateModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cover` | `string` | No |  |
| `dependency_id` | `string` | No |  |
| `description` | `string` | No |  |
| `price` | `number` | No |  |
| `thumbnail` | `string` | No |  |
| `visibility` | `string` | No |  |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Model` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestUpdateModelRequest{
    Cover: "...",  // string
    DependencyID: "...",  // string
    Description: "...",  // string
    Price: "...",  // number
    Thumbnail: "...",  // string
}
resp, err := client.Models().Model.PutModelByID(model.NewPutModelByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDInfo`

**`GET /api-key/model/{id}/info`** — Get Api Key Model Info

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.ApiKeyInfoResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ApiKeyInfo` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ApiKeyInfo`**

| Field | Type | Description |
| --- | --- | --- |
| `api_price` | `number` |  |
| `commit_hash` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_description` | `string` |  |
| `model_id` | `string` |  |
| `output_format` | `map[string]any` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelByIDInfo(model.NewGetModelByIDInfoParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutModelByIDMetadata`

**`PUT /api-key/model/{id}/metadata`** — Update Model Metadata

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.UpdateModelMetadataRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `language` | `array[string]` | No | en, vi |
| `library` | `array[string]` | No | tf, jax |
| `license` | `string` | No |  |
| `pretty_name` | `string` | No |  |
| `task` | `string` | No |  |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Model` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.RequestUpdateModelMetadataRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    PrettyName: "...",  // string
    Task: "...",  // string
}
resp, err := client.Models().Model.PutModelByIDMetadata(model.NewPutModelByIDMetadataParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDServing`

**`GET /api-key/model/{id}/serving`** — Check Model Is Serving

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.CheckModelIsServingResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckModelIsServingData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckModelIsServingData`**

| Field | Type | Description |
| --- | --- | --- |
| `consumers` | `integer` |  |
| `serving` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelByIDServing(model.NewGetModelByIDServingParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDStatistics`

**`POST /api-key/model/{id}/statistics`** — Get Model Statistics

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

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
resp, err := client.Models().Model.PostModelByIDStatistics(model.NewPostModelByIDStatisticsParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDTaskCost`

**`GET /api-key/model/{id}/task/cost`** — Get cost to compute task by model api key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

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
resp, err := client.Models().Model.GetModelByIDTaskCost(model.NewGetModelByIDTaskCostParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDVersioning`

**`GET /api-key/model/{id}/versioning`** — Get Current Model Versioning By Model Id By ApiKey

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.ModelVersioningGroupLiteResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ModelVersioningGroupLite` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ModelVersioningGroupLite`**

| Field | Type | Description |
| --- | --- | --- |
| `commit_hash` | `string` | Dependency        map[string]interface{} `json:"dependency"` |
| `commit_message` | `string` | TestResult        map[string]interface{} `json:"test_result"`
InputFormat       map[string]interface{} `json:"input_format"`
OutputFormat      map[string]interface{} `json:"output_format"`
SysRequired       map[string]interface{} `json:"sys_require"` |
| `is_active` | `boolean` |  |
| `last_checked_at` | `string` |  |
| `model_id` | `string` |  |
| `pending_platforms` | `array[models.PlatformTask]` |  |
| `rejected_platforms` | `array[models.PlatformTask]` |  |
| `verified_platforms` | `array[models.PlatformTask]` |  |
| `verify_status` | `string` |  |

**`models.PlatformTask`**

| Field | Type | Description |
| --- | --- | --- |
| `platform` | `string` |  |
| `test_result` | `map[string]any` |  |
| `updated_at` | `string` |  |
| `verify_task_id` | `string` |  |

**`models.PlatformTask`**

| Field | Type | Description |
| --- | --- | --- |
| `platform` | `string` |  |
| `test_result` | `map[string]any` |  |
| `updated_at` | `string` |  |
| `verify_task_id` | `string` |  |

**`models.PlatformTask`**

| Field | Type | Description |
| --- | --- | --- |
| `platform` | `string` |  |
| `test_result` | `map[string]any` |  |
| `updated_at` | `string` |  |
| `verify_task_id` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelByIDVersioning(model.NewGetModelByIDVersioningParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByUsernameByName`

**`GET /api-key/model/{username}/{name}`** — Get Model By Name

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Model's username |
| `name` | path | `string` | Yes | Model's name |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Model` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models().Model.GetModelByUsernameByName(model.NewGetModelByUsernameByNameParams().WithContext(ctx).WithUsername("<username>").WithName("<name>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
