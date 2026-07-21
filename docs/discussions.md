# Discussions — Go SDK

> Discussion forums — create and manage discussions. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

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
req := &models.RequestUpdateCommentRequest{
    Content: "...",  // string
    TagUsernames: "...",  // array[string]
}
resp, err := client.Discussions().Comment.PutCommentsByID(comment.NewPutCommentsByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
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

**200 OK** — `response.LikeResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.IsLikedByUser` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.IsLikedByUser`**

| Field | Type | Description |
| --- | --- | --- |
| `is_liked_by_user` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Discussions().Reaction.PutDatasetByIDLike(reaction.NewPutDatasetByIDLikeParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetDiscussionCompetitionByID`

**`GET /api-key/discussion/competition/{id}`** — Get competition discussion list

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Competition's id |
| `filter` | query | `string` | No |  |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `search` | query | `string` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.DiscussionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DiscussionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DiscussionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteDiscussion]` |  |
| `total` | `integer` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
resp, err := client.Discussions().Discussion.GetDiscussionCompetitionByID(discussion.NewGetDiscussionCompetitionByIDParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDiscussionCompetitionByID`

**`POST /api-key/discussion/competition/{id}`** — Create competition discussion

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Competition's id |

**Request Body** — `request.CreateCompetitionDiscussionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes |  |
| `title` | `string` | Yes |  |

**Responses**

**200 OK** — `response.DiscussionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteDiscussion` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
req := &models.RequestCreateCompetitionDiscussionRequest{
    Content: "...",  // string  // required
    Title: "...",  // string  // required
}
resp, err := client.Discussions().Discussion.PostDiscussionCompetitionByID(discussion.NewPostDiscussionCompetitionByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetDiscussionDatasetByID`

**`GET /api-key/discussion/dataset/{id}`** — Get dataset discussion list

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Dataset's id |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.DiscussionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DiscussionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DiscussionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteDiscussion]` |  |
| `total` | `integer` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
resp, err := client.Discussions().Discussion.GetDiscussionDatasetByID(discussion.NewGetDiscussionDatasetByIDParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDiscussionDatasetByID`

**`POST /api-key/discussion/dataset/{id}`** — Create dataset discussion

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Dataset's id |

**Request Body** — `request.CreateDatasetDiscussionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes |  |
| `title` | `string` | Yes |  |

**Responses**

**200 OK** — `response.DiscussionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteDiscussion` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
req := &models.RequestCreateDatasetDiscussionRequest{
    Content: "...",  // string  // required
    Title: "...",  // string  // required
}
resp, err := client.Discussions().Discussion.PostDiscussionDatasetByID(discussion.NewPostDiscussionDatasetByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetDiscussionModelByID`

**`GET /api-key/discussion/model/{id}`** — Get model discussion list

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.DiscussionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DiscussionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DiscussionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteDiscussion]` |  |
| `total` | `integer` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
resp, err := client.Discussions().Discussion.GetDiscussionModelByID(discussion.NewGetDiscussionModelByIDParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostDiscussionModelByID`

**`POST /api-key/discussion/model/{id}`** — Create model discussion

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CreateModelDiscussionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | Yes |  |
| `title` | `string` | Yes |  |

**Responses**

**200 OK** — `response.DiscussionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteDiscussion` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
req := &models.RequestCreateModelDiscussionRequest{
    Content: "...",  // string  // required
    Title: "...",  // string  // required
}
resp, err := client.Discussions().Discussion.PostDiscussionModelByID(discussion.NewPostDiscussionModelByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutDiscussionByID`

**`PUT /api-key/discussion/{id}`** — Update discussion

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Discussion's id |

**Request Body** — `request.UpdateDiscussionRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `content` | `string` | No |  |
| `is_closed` | `boolean` | No |  |
| `title` | `string` | No |  |

**Responses**

**200 OK** — `response.DiscussionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteDiscussion` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteDiscussion`**

| Field | Type | Description |
| --- | --- | --- |
| `comments_count` | `integer` |  |
| `content` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `is_closed` | `boolean` |  |
| `owner` | `models.Owner` |  |
| `reacted` | `models.Reaction` |  |
| `reactions_statistics` | `array[models.ReactionStats]` |  |
| `reacts_count` | `integer` |  |
| `title` | `string` |  |
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
req := &models.RequestUpdateDiscussionRequest{
    Content: "...",  // string
    IsClosed: "...",  // boolean
    Title: "...",  // string
}
resp, err := client.Discussions().Discussion.PutDiscussionByID(discussion.NewPutDiscussionByIDParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
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
resp, err := client.Discussions().Comment.GetDiscussionByIDComments(comment.NewGetDiscussionByIDCommentsParams().WithContext(ctx).WithID("<id>"), nil)
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
req := &models.RequestCreateCommentRequest{
    Content: "...",  // string
    TagUsernames: "...",  // array[string]
}
resp, err := client.Discussions().Comment.PostDiscussionByIDComments(comment.NewPostDiscussionByIDCommentsParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
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
req := &models.RequestReactItemRequest{
    Itemname: "...",  // string
    Reactname: "...",  // string
}
resp, err := client.Discussions().Reaction.PutItemsByIDReact(reaction.NewPutItemsByIDReactParams().WithContext(ctx).WithID("<id>").WithInput(req), nil)
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

**200 OK** — `response.LikeResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.IsLikedByUser` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.IsLikedByUser`**

| Field | Type | Description |
| --- | --- | --- |
| `is_liked_by_user` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Discussions().Reaction.PutModelByIDLike(reaction.NewPutModelByIDLikeParams().WithContext(ctx).WithID("<id>"), nil)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
