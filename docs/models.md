# Models — Go SDK

> AI model management — create, train, deploy, and version models. Requires: `x-api-key` header.

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
| `cover` | `string` | No |  |
| `dependency_id` | `string` | No |  |
| `description` | `string` | No |  |
| `language` | `array[string]` | No | en, vi |
| `library` | `array[string]` | No | tf |
| `license` | `string` | Yes |  |
| `name` | `string` | Yes |  |
| `pretty_name` | `string` | No |  |
| `price` | `number` | No |  |
| `task` | `string` | No |  |
| `thumbnail` | `string` | No |  |
| `visibility` | `string` | Yes | One of: `public`, `private` |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
req := &models.CreateModelRequest{
    AuthorID: "...",  // string
    Cover: "...",  // string
    DependencyID: "...",  // string
    Description: "...",  // string
    Language: "...",  // array[string]
}
resp, err := client.Models.Model.PostModel(ctx, req)
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
req := &models.GetModelListRequest{
    FilterBy: "...",  // string
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
}
resp, err := client.Models.Model.PostModelList(ctx, req)
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
req := &models.GetModelListByAuthorRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Models.Model.PostModelListByAuthorByUsername(ctx, req)
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

**200 OK** — `response.ListDataResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array[string]` |  |
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
req := &models.MatchingModelsTagsRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    TagType: "...",  // string
    Task: "...",  // string
}
resp, err := client.Models.Model.PostModelMatchingTags(ctx, req)
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
resp, err := client.Models.Model.GetModelOrganizationByOrg(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteModelTaskReviewsByID`

**`DELETE /api-key/model/task/reviews/{id}`** — Delete Task Reviews By Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Task Review's id |

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
resp, err := client.Models.Reviews.DeleteModelTaskReviewsByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelTaskByIDReviews`

**`GET /api-key/model/task/{id}/reviews`** — Get Task Review By Task Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Task's id |

**Request Body** — `request.GetTaskReviewByTaskIdRequest`

_No fields defined._

**Responses**

**200 OK** — `response.TaskReviewsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.TaskReviews` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.TaskReviews`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `model_id` | `string` |  |
| `owner` | `models.Owner` |  |
| `point` | `integer` |  |
| `result` | `map[string]any` |  |
| `task_id` | `string` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetTaskReviewByTaskIdRequest{
    
}
resp, err := client.Models.Reviews.GetModelTaskByIDReviews(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelTrainingCost`

**`POST /api-key/model/training/cost`** — Calculate Cost To Training Ai Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.CalculateCostToTrainingAiModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `dataset_id` | `string` | No |  |
| `training_task_id` | `string` | No |  |

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
req := &models.CalculateCostToTrainingAiModelRequest{
    DatasetID: "...",  // string
    TrainingTaskID: "...",  // string
}
resp, err := client.Models.Training.PostModelTrainingCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelTrainingTask`

**`GET /api-key/model/training/task`** — Get List User Training Task

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.GetListUserTrainingTaskRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |

**Responses**

**200 OK** — `response.UserTrainingTaskListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.UserTrainingTaskListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.UserTrainingTaskListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.UserTrainingTask]` |  |
| `total` | `integer` |  |

**`models.UserTrainingTask`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `dataset_id` | `string` |  |
| `dataset_repo_size` | `number` |  |
| `dataset_repo_url` | `string` |  |
| `dataset_sha` | `string` |  |
| `id` | `string` |  |
| `task_id` | `string` |  |
| `training_data` | `map[string]any` |  |
| `training_metadata` | `map[string]any` |  |
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
req := &models.GetListUserTrainingTaskRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Models.Training.GetModelTrainingTask(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteModelTrainingTaskByID`

**`DELETE /api-key/model/training/task/{id}`** — Delete User Training Task By Task Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Training Task's id |

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
resp, err := client.Models.Training.DeleteModelTrainingTaskByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelVerifyHubTaskByID`

**`GET /api-key/model/verify/hub/task/{id}`** — Get Model Versioning By Hub Task Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Hub Task's id |

**Responses**

**200 OK** — `response.ModelVersioningResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ModelVersioning` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ModelVersioning`**

| Field | Type | Description |
| --- | --- | --- |
| `commit_hash` | `string` |  |
| `commit_message` | `string` |  |
| `created_at` | `string` |  |
| `dependency` | `map[string]any` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `node_reward` | `number` |  |
| `output_format` | `map[string]any` |  |
| `platform` | `string` |  |
| `sys_require` | `map[string]any` |  |
| `test_result` | `map[string]any` |  |
| `updated_at` | `string` |  |
| `verify_status` | `string` |  |
| `verify_task_id` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models.Verify.GetModelVerifyHubTaskByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelVerifySupportPlatforms`

**`GET /api-key/model/verify/support/platforms`** — Get List Platforms Support

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.GetListPlatformSupportResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array[string]` |  |
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
resp, err := client.Models.Model.GetModelVerifySupportPlatforms(ctx)
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
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
resp, err := client.Models.Model.GetModelByID(ctx, "<id>")
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
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
req := &models.UpdateModelRequest{
    Cover: "...",  // string
    DependencyID: "...",  // string
    Description: "...",  // string
    Price: "...",  // number
    Thumbnail: "...",  // string
}
resp, err := client.Models.Model.PutModelByID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteModelByID`

**`DELETE /api-key/model/{id}`** — Delete Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.DeleteModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `name` | `string` | Yes |  |

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
req := &models.DeleteModelRequest{
    Name: "...",  // string  // required
}
resp, err := client.Models.Model.DeleteModelByID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDApiKey`

**`GET /api-key/model/{id}/api-key`** — Get List Model ApiKey

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

**Responses**

**200 OK** — `response.ListApiKeyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ListApiKeyData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ListApiKeyData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteApiKey]` |  |
| `total` | `integer` |  |

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
resp, err := client.Models.Model.GetModelByIDApiKey(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDApiKey`

**`POST /api-key/model/{id}/api-key`** — Create Model ApiKey

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CreateApiKeyRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `boolean` | No |  |
| `description` | `string` | No |  |
| `name` | `string` | Yes |  |
| `org_username` | `string` | No |  |

**Responses**

**200 OK** — `response.ApiKeyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ApiKey` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ApiKey`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `api_key` | `string` |  |
| `api_price` | `number` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `name` | `string` |  |
| `output_format` | `map[string]any` |  |
| `total_cost` | `number` |  |
| `total_failed` | `integer` |  |
| `total_request` | `integer` |  |
| `total_success` | `integer` |  |
| `updated_at` | `string` |  |
| `user` | `models.User` |  |
| `user_id` | `string` |  |

**`models.User`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_request_to_join` | `boolean` |  |
| `avatar_url` | `string` |  |
| `bio` | `string` |  |
| `blocked` | `boolean` |  |
| `blocked_at` | `string` |  |
| `email` | `string` |  |
| `followers` | `array[models.Follow]` |  |
| `followers_count` | `integer` |  |
| `followings` | `array[models.Follow]` |  |
| `followings_count` | `integer` |  |
| `github_link` | `string` |  |
| `github_name` | `string` |  |
| `home_page_name` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `invite_offers` | `array[models.Offer]` |  |
| `invite_offers_count` | `integer` |  |
| `is_following` | `boolean` |  |
| `join_id` | `string` |  |
| `join_offers` | `array[models.Offer]` |  |
| `join_offers_count` | `integer` |  |
| `members` | `array[models.Member]` |  |
| `members_count` | `integer` |  |
| `name` | `string` |  |
| `role` | `string` |  |
| `token` | `string` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |
| `wallet` | `models.Wallet` |  |
| `wallet_address` | `string` |  |
| `wallet_connection` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

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

**`models.Member`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `full_name` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Wallet`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `string` |  |
| `debt` | `string` |  |
| `earnings` | `string` |  |
| `free_balance` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CreateApiKeyRequest{
    Active: "...",  // boolean
    Description: "...",  // string
    Name: "...",  // string  // required
    OrgUsername: "...",  // string
}
resp, err := client.Models.Model.PostModelByIDApiKey(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDDownload`

**`GET /api-key/model/{id}/download`** — Get List Model Download

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

**Responses**

**200 OK** — `response.DownloadModelListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DownloadModelListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DownloadModelListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteDownloadModel]` |  |
| `total` | `integer` |  |

**`models.LiteDownloadModel`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models.Model.GetModelByIDDownload(ctx, "<id>")
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
resp, err := client.Models.Model.GetModelByIDInfo(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDLike`

**`GET /api-key/model/{id}/like`** — Get List Model Like

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

**Responses**

**200 OK** — `response.LikeModelListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.LikeModelListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.LikeModelListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.LiteLikeModel]` |  |
| `total` | `integer` |  |

**`models.LiteLikeModel`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `owner` | `models.Owner` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models.Model.GetModelByIDLike(ctx, "<id>")
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
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
req := &models.UpdateModelMetadataRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    PrettyName: "...",  // string
    Task: "...",  // string
}
resp, err := client.Models.Model.PutModelByIDMetadata(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDPreVerify`

**`POST /api-key/model/{id}/pre-verify`** — Check Valid Source code To Verify Ai Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CheckValidToVerifyAiModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `commit_hash` | `string` | Yes |  |
| `platforms` | `array[string]` | Yes |  |

**Responses**

**200 OK** — `response.CheckValidToVerifyAiModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckValidToVerifyAiModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckValidToVerifyAiModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `valid` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CheckValidToVerifyAiModelRequest{
    CommitHash: "...",  // string  // required
    Platforms: "...",  // array[string]  // required
}
resp, err := client.Models.Verify.PostModelByIDPreVerify(ctx, req)
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
resp, err := client.Models.Model.GetModelByIDServing(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDSetting`

**`GET /api-key/model/{id}/setting`** — Get Model Setting By Model Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.ModelSettingResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ModelSetting` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ModelSetting`**

| Field | Type | Description |
| --- | --- | --- |
| `api_price` | `number` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `output_format` | `map[string]any` |  |
| `sys_req_cpu_cores` | `integer` |  |
| `sys_req_gpu_memory` | `integer` |  |
| `sys_req_ram` | `integer` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models.Setting.GetModelByIDSetting(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutModelByIDSetting`

**`PUT /api-key/model/{id}/setting`** — Update Model Setting

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.UpdateModelSettingRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `api_price` | `number` | No |  |
| `sys_req_cpu_cores` | `integer` | No |  |
| `sys_req_gpu_memory` | `integer` | No |  |
| `sys_req_ram` | `integer` | No |  |

**Responses**

**200 OK** — `response.ModelSettingResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ModelSetting` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ModelSetting`**

| Field | Type | Description |
| --- | --- | --- |
| `api_price` | `number` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `output_format` | `map[string]any` |  |
| `sys_req_cpu_cores` | `integer` |  |
| `sys_req_gpu_memory` | `integer` |  |
| `sys_req_ram` | `integer` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.UpdateModelSettingRequest{
    APIPrice: "...",  // number
    SysReqCpuCores: "...",  // integer
    SysReqGpuMemory: "...",  // integer
    SysReqRam: "...",  // integer
}
resp, err := client.Models.Setting.PutModelByIDSetting(ctx, req)
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
resp, err := client.Models.Model.PostModelByIDStatistics(ctx, req)
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
resp, err := client.Models.Model.GetModelByIDTaskCost(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDTaskReviews`

**`POST /api-key/model/{id}/task/reviews`** — Create Task Reviews

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CreateTaskReviewsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `description` | `string` | No |  |
| `point` | `integer` | No |  |
| `task_id` | `string` | No |  |

**Responses**

**200 OK** — `response.TaskReviewsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.TaskReviews` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.TaskReviews`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `model_id` | `string` |  |
| `owner` | `models.Owner` |  |
| `point` | `integer` |  |
| `result` | `map[string]any` |  |
| `task_id` | `string` |  |
| `updated_at` | `string` |  |

**`models.Owner`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CreateTaskReviewsRequest{
    Description: "...",  // string
    Point: "...",  // integer
    TaskID: "...",  // string
}
resp, err := client.Models.Reviews.PostModelByIDTaskReviews(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDVerify`

**`POST /api-key/model/{id}/verify`** — Verify Ai Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.VerifyAiModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `commit_hash` | `string` | No |  |
| `platforms` | `array[string]` | Yes |  |

**Responses**

**200 OK** — `response.VerifyAiModelResponse`

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
req := &models.VerifyAiModelRequest{
    CommitHash: "...",  // string
    Platforms: "...",  // array[string]  // required
}
resp, err := client.Models.Verify.PostModelByIDVerify(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostModelByIDVerifyCost`

**`POST /api-key/model/{id}/verify/cost`** — Calculate Cost To Verify Ai Model

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CalculateCostToVerifyAiModelRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `commit_hash` | `string` | Yes |  |
| `platforms` | `array[string]` | Yes |  |

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
req := &models.CalculateCostToVerifyAiModelRequest{
    CommitHash: "...",  // string  // required
    Platforms: "...",  // array[string]  // required
}
resp, err := client.Models.Verify.PostModelByIDVerifyCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDVerifyPending`

**`GET /api-key/model/{id}/verify/pending`** — Check Model Pending State

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Request Body** — `request.CheckModelPendingStateRequest`

_No fields defined._

**Responses**

**200 OK** — `response.CheckModelPendingStateResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckModelPendingStateData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckModelPendingStateData`**

| Field | Type | Description |
| --- | --- | --- |
| `pending_state` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CheckModelPendingStateRequest{
    
}
resp, err := client.Models.Verify.GetModelByIDVerifyPending(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDVerifyTask`

**`GET /api-key/model/{id}/verify/task`** — Get List Verify Model Task By Commit Hash And Status

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |
| `commitHash` | query | `string` | Yes |  |
| `verifyStatus` | query | `string` | No |  |

**Responses**

**200 OK** — `response.ModelVersioningGroupLiteListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelVersioningGroupLiteListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelVersioningGroupLiteListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.ModelVersioningGroupLite]` |  |
| `total` | `integer` |  |

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
resp, err := client.Models.Verify.GetModelByIDVerifyTask(ctx, "<id>")
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
resp, err := client.Models.Versioning.GetModelByIDVersioning(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutModelByIDVersioning`

**`PUT /api-key/model/{id}/versioning`** — Change Model Versioning By Commit Hash

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |
| `commitHash` | query | `string` | Yes |  |

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
resp, err := client.Models.Versioning.PutModelByIDVersioning(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteModelByIDVersioning`

**`DELETE /api-key/model/{id}/versioning`** — Delete Model Versioning By Commit Hash

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |
| `commitHash` | query | `string` | Yes |  |

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
resp, err := client.Models.Versioning.DeleteModelByIDVersioning(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetModelByIDVersioningList`

**`GET /api-key/model/{id}/versioning/list`** — Get Verified List Model Versioning

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
| `verifyStatus` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetListModelVersioningLiteResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetListModelVersioningsLiteData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetListModelVersioningsLiteData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.ModelVersioningGroupLite]` |  |
| `total` | `integer` |  |

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
resp, err := client.Models.Versioning.GetModelByIDVersioningList(ctx, "<id>")
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
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
resp, err := client.Models.Model.GetModelByUsernameByName(ctx, "<username>", "<name>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPackageList`

**`POST /api-key/package/list`** — Get Api Package List

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.GetApiPackageListRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |

**Responses**

**200 OK** — `response.ApiPackageListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ApiPackageListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ApiPackageListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.ApiPackage]` |  |
| `total` | `integer` |  |

**`models.ApiPackage`**

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
req := &models.GetApiPackageListRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Models.Package.PostPackageList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetUserApiKeyByID`

**`GET /api-key/user/api-key/{id}`** — Get Api Key Detail By Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Api Key's id |

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
resp, err := client.Models.APIKey.GetUserApiKeyByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutUserApiKeyByID`

**`PUT /api-key/user/api-key/{id}`** — Update Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Api Key's id |

**Request Body** — `request.UpdateApiKeyRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `active` | `boolean` | No |  |
| `description` | `string` | No |  |
| `name` | `string` | No |  |
| `org_username` | `string` | No |  |

**Responses**

**200 OK** — `response.ApiKeyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ApiKey` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ApiKey`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `api_key` | `string` |  |
| `api_price` | `number` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `name` | `string` |  |
| `output_format` | `map[string]any` |  |
| `total_cost` | `number` |  |
| `total_failed` | `integer` |  |
| `total_request` | `integer` |  |
| `total_success` | `integer` |  |
| `updated_at` | `string` |  |
| `user` | `models.User` |  |
| `user_id` | `string` |  |

**`models.User`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_request_to_join` | `boolean` |  |
| `avatar_url` | `string` |  |
| `bio` | `string` |  |
| `blocked` | `boolean` |  |
| `blocked_at` | `string` |  |
| `email` | `string` |  |
| `followers` | `array[models.Follow]` |  |
| `followers_count` | `integer` |  |
| `followings` | `array[models.Follow]` |  |
| `followings_count` | `integer` |  |
| `github_link` | `string` |  |
| `github_name` | `string` |  |
| `home_page_name` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `invite_offers` | `array[models.Offer]` |  |
| `invite_offers_count` | `integer` |  |
| `is_following` | `boolean` |  |
| `join_id` | `string` |  |
| `join_offers` | `array[models.Offer]` |  |
| `join_offers_count` | `integer` |  |
| `members` | `array[models.Member]` |  |
| `members_count` | `integer` |  |
| `name` | `string` |  |
| `role` | `string` |  |
| `token` | `string` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |
| `wallet` | `models.Wallet` |  |
| `wallet_address` | `string` |  |
| `wallet_connection` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

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

**`models.Member`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `full_name` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Wallet`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `string` |  |
| `debt` | `string` |  |
| `earnings` | `string` |  |
| `free_balance` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.UpdateApiKeyRequest{
    Active: "...",  // boolean
    Description: "...",  // string
    Name: "...",  // string
    OrgUsername: "...",  // string
}
resp, err := client.Models.APIKey.PutUserApiKeyByID(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostUserApiKeyByIDStatistics`

**`POST /api-key/user/api-key/{id}/statistics`** — Get Api Key By Id

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Api Key's id |

**Request Body** — `request.GetApiKeyByIdRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `org_username` | `string` | No |  |
| `to` | `string` | No | 2023-05-07 15:04:05 |

**Responses**

**200 OK** — `response.ApiKeyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.ApiKey` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.ApiKey`**

| Field | Type | Description |
| --- | --- | --- |
| `active` | `boolean` |  |
| `api_key` | `string` |  |
| `api_price` | `number` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `input_format` | `map[string]any` |  |
| `model_id` | `string` |  |
| `name` | `string` |  |
| `output_format` | `map[string]any` |  |
| `total_cost` | `number` |  |
| `total_failed` | `integer` |  |
| `total_request` | `integer` |  |
| `total_success` | `integer` |  |
| `updated_at` | `string` |  |
| `user` | `models.User` |  |
| `user_id` | `string` |  |

**`models.User`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_request_to_join` | `boolean` |  |
| `avatar_url` | `string` |  |
| `bio` | `string` |  |
| `blocked` | `boolean` |  |
| `blocked_at` | `string` |  |
| `email` | `string` |  |
| `followers` | `array[models.Follow]` |  |
| `followers_count` | `integer` |  |
| `followings` | `array[models.Follow]` |  |
| `followings_count` | `integer` |  |
| `github_link` | `string` |  |
| `github_name` | `string` |  |
| `home_page_name` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `invite_offers` | `array[models.Offer]` |  |
| `invite_offers_count` | `integer` |  |
| `is_following` | `boolean` |  |
| `join_id` | `string` |  |
| `join_offers` | `array[models.Offer]` |  |
| `join_offers_count` | `integer` |  |
| `members` | `array[models.Member]` |  |
| `members_count` | `integer` |  |
| `name` | `string` |  |
| `role` | `string` |  |
| `token` | `string` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |
| `wallet` | `models.Wallet` |  |
| `wallet_address` | `string` |  |
| `wallet_connection` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

**`models.Follow`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

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

**`models.Member`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar_url` | `string` |  |
| `full_name` | `string` |  |
| `id` | `string` |  |
| `username` | `string` |  |

**`models.Wallet`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `string` |  |
| `debt` | `string` |  |
| `earnings` | `string` |  |
| `free_balance` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetApiKeyByIdRequest{
    From: "...",  // string
    OrgUsername: "...",  // string
    To: "...",  // string
}
resp, err := client.Models.APIKey.PostUserApiKeyByIDStatistics(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetUserPlaygroundRemaining`

**`GET /api-key/user/playground/remaining`** — Get Playground Uses Remaining By User

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Responses**

**200 OK** — `response.PlaygroundPackageListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.PlaygroundPackageListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.PlaygroundPackageListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.PlaygroundPackage]` |  |
| `total` | `integer` |  |

**`models.PlaygroundPackage`**

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
resp, err := client.Models.Playground.GetUserPlaygroundRemaining(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicModelList`

**`POST /public/model/list`** — Get Public Model List

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
req := &models.GetModelListRequest{
    FilterBy: "...",  // string
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
}
resp, err := client.Models.Model.PostPublicModelList(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicModelListByAuthorByUsername`

**`POST /public/model/list-by-author/{username}`** — Get Public Model List By User

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Username |

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
req := &models.GetModelListByAuthorRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Models.Model.PostPublicModelListByAuthorByUsername(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicModelMatchingTags`

**`POST /public/model/matching-tags`** — Matching Public Models Tags

**Request Body** — `request.MatchingModelsTagsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `language` | `array[string]` | No |  |
| `library` | `array[string]` | No |  |
| `license` | `string` | No |  |
| `tag_type` | `string` | No |  |
| `task` | `string` | No |  |

**Responses**

**200 OK** — `response.ListDataResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array[string]` |  |
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
req := &models.MatchingModelsTagsRequest{
    Language: "...",  // array[string]
    Library: "...",  // array[string]
    License: "...",  // string
    TagType: "...",  // string
    Task: "...",  // string
}
resp, err := client.Models.Model.PostPublicModelMatchingTags(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelMetadata`

**`GET /public/model/metadata`** — Get Model Metadata

**Responses**

**200 OK** — `response.MetadataListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.MetadataListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.MetadataListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Metadata]` |  |

**`models.Metadata`**

| Field | Type | Description |
| --- | --- | --- |
| `category` | `string` |  |
| `id` | `string` |  |
| `label` | `string` |  |
| `type` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Models.Model.GetPublicModelMetadata(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelOrganizationByOrg`

**`GET /public/model/organization/{org}`** — Get List Model By Org Username

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
resp, err := client.Models.Model.GetPublicModelOrganizationByOrg(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelTrending`

**`GET /public/model/trending`** — Get List Models Trending

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |

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
resp, err := client.Models.Model.GetPublicModelTrending(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelByID`

**`GET /public/model/{id}`** — Get public model by id

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `id` | path | `string` | Yes | Model's id |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
resp, err := client.Models.Model.GetPublicModelByID(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelByIDVersioning`

**`GET /public/model/{id}/versioning`** — Get Current Model Versioning By Model Id

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
resp, err := client.Models.Model.GetPublicModelByIDVersioning(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicModelByUsernameByName`

**`GET /public/model/{username}/{name}`** — Get Public Model By Name

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Username |
| `name` | path | `string` | Yes | Model's name |

**Responses**

**200 OK** — `response.ModelResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModelData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModelData`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `downloads` | `integer` |  |
| `id` | `string` |  |
| `is_liked_by_user` | `boolean` |  |
| `likes` | `integer` |  |
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
resp, err := client.Models.Model.GetPublicModelByUsernameByName(ctx, "<username>", "<name>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
