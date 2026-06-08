# Collections — Go SDK

> Collection management — organize and curate resources. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `PostCollection`

**`POST /api-key/collection`** — Create Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.CreateCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `thumbnail` | `string` | No |  |
| `visibility` | `string` | Yes | One of: `public`, `private` |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.CreateCollectionRequest{
    Description: "...",  // string  // required
    Name: "...",  // string  // required
    Thumbnail: "...",  // string
    Visibility: "...",  // string  // required
}
resp, err := client.Collections.Collection.PostCollection(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostCollectionList`

**`POST /api-key/collection/list`** — Get List Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.GetListCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `string` | No | One of: `public`, `community`, `author` |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `desc`, `asc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `votes`, `updated`, `oldest` |

**Responses**

**200 OK** — `response.CollectionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CollectionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CollectionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Collection]` |  |
| `total` | `integer` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.GetListCollectionRequest{
    Filter: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
    Order: "...",  // string
    Search: "...",  // string
}
resp, err := client.Collections.Collection.PostCollectionList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetCollectionByID`

**`GET /api-key/collection/{id}`** — Get Collection By Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
resp, err := client.Collections.Collection.GetCollectionByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutCollectionByID`

**`PUT /api-key/collection/{id}`** — Update Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Request Body** — `request.UpdateCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | No |  |
| `name` | `string` | No |  |
| `thumbnail` | `string` | No |  |
| `visibility` | `string` | No | One of: `public`, `private` |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.UpdateCollectionRequest{
    Description: "...",  // string
    Name: "...",  // string
    Thumbnail: "...",  // string
    Visibility: "...",  // string
}
resp, err := client.Collections.Collection.PutCollectionByID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteCollectionByID`

**`DELETE /api-key/collection/{id}`** — Delete Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
resp, err := client.Collections.Collection.DeleteCollectionByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostCollectionByIDItem`

**`POST /api-key/collection/{id}/item`** — Add Item To Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Request Body** — `request.AddItemToCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `item_id` | `string` | No |  |
| `type` | `string` | Yes | One of: `dataset`, `model` |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.AddItemToCollectionRequest{
    ItemID: "...",  // string
    Type: "...",  // string  // required
}
resp, err := client.Collections.Collection.PostCollectionByIDItem(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteCollectionByIDItem`

**`DELETE /api-key/collection/{id}/item`** — Remove Item From Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Request Body** — `request.RemoveItemFromCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `item_id` | `string` | Yes |  |
| `type` | `string` | Yes | One of: `dataset`, `model` |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.RemoveItemFromCollectionRequest{
    ItemID: "...",  // string  // required
    Type: "...",  // string  // required
}
resp, err := client.Collections.Collection.DeleteCollectionByIDItem(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutCollectionByIDVote`

**`PUT /api-key/collection/{id}/vote`** — Vote Collection

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
resp, err := client.Collections.Collection.PutCollectionByIDVote(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostCollectionByUsernameList`

**`POST /api-key/collection/{username}/list`** — Get Collection List By Username

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Username |

**Request Body** — `request.GetListCollectionByUsernameRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `desc`, `asc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `votes`, `updated`, `oldest` |

**Responses**

**200 OK** — `response.CollectionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CollectionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CollectionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Collection]` |  |
| `total` | `integer` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.GetListCollectionByUsernameRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
    Order: "...",  // string
    Search: "...",  // string
    Sort: "...",  // string
}
resp, err := client.Collections.Collection.PostCollectionByUsernameList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicCollectionList`

**`POST /public/collection/list`** — Get List Public Collection

**Request Body** — `request.GetListCollectionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `filter` | `string` | No | One of: `public`, `community`, `author` |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `desc`, `asc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `votes`, `updated`, `oldest` |

**Responses**

**200 OK** — `response.CollectionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CollectionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CollectionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Collection]` |  |
| `total` | `integer` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.GetListCollectionRequest{
    Filter: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
    Order: "...",  // string
    Search: "...",  // string
}
resp, err := client.Collections.Collection.PostPublicCollectionList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicCollectionByID`

**`GET /public/collection/{id}`** — Get Public Collection By Id

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Collection's id |

**Responses**

**200 OK** — `response.CollectionReponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.Collection` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
resp, err := client.Collections.Collection.GetPublicCollectionByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicCollectionByUsernameList`

**`POST /public/collection/{username}/list`** — Get List Public Collections By Username

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | username |

**Request Body** — `request.GetListCollectionByUsernameRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `order` | `string` | No | One of: `desc`, `asc` |
| `search` | `string` | No |  |
| `sort` | `string` | No | One of: `trending`, `votes`, `updated`, `oldest` |

**Responses**

**200 OK** — `response.CollectionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CollectionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CollectionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Collection]` |  |
| `total` | `integer` |  |

**`models.Collection`**

| Field | Type | Description |
| --- | --- | --- |
| `author_avatar` | `string` |  |
| `created_at` | `string` |  |
| `datasets` | `array[models.Dataset]` |  |
| `datasets_count` | `integer` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `is_voted_by_user` | `boolean` |  |
| `models` | `array[models.Model]` |  |
| `models_count` | `integer` |  |
| `name` | `string` |  |
| `thumbnail` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `violated` | `boolean` |  |
| `visibility` | `string` |  |
| `votes_count` | `integer` |  |

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
req := &models.GetListCollectionByUsernameRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
    Order: "...",  // string
    Search: "...",  // string
    Sort: "...",  // string
}
resp, err := client.Collections.Collection.PostPublicCollectionByUsernameList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
