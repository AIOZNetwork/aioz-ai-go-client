# Public — Go SDK

> Public endpoints — health checks, public listings, and medals. Public — no authentication required.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetPublicContentTermsPrivacy`

**`GET /public/content/terms-privacy`** — Get Content Config By Static Keys

**Responses**

**200 OK** — `response.TermsAndPrivacyResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array[object]` |  |
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
resp, err := client.Publics.Public.GetPublicContentTermsPrivacy(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicContentByKeys`

**`GET /public/content/{keys}`** — Get Content Config By List Keys

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `keys` | path | `string` | Yes | List Keys |

**Responses**

**200 OK** — `response.ContentConfigListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ContentConfigListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ContentConfigListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.ContentConfig]` |  |
| `total` | `integer` |  |

**`models.ContentConfig`**

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `key` | `string` |  |
| `updated_at` | `string` |  |
| `value` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicContentByKeys(ctx, "<keys>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostPublicCostEstimating`

**`POST /public/cost/estimating`** — Estimate Cost

**Request Body** — `request.EstimateCostRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `total_size` | `number` | Yes |  |

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
req := &models.EstimateCostRequest{
    TotalSize: "...",  // number  // required
}
resp, err := client.Publics.Public.PostPublicCostEstimating(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicMetadata`

**`GET /public/metadata`** — GetListMetadata

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `category` | query | `string` | No |  |
| `types` | query | `array` | No |  |

**Responses**

**200 OK** — `response.GetListMetadataResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array[models.MetadataListByType]` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.MetadataListByType`**

| Field | Type | Description |
| --- | --- | --- |
| `items` | `array[models.LiteMetadata]` |  |
| `type` | `string` |  |

**`models.LiteMetadata`**

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `label` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicMetadata(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicRepositoryByOwnerusernameByRepositorynameContentReadme`

**`GET /public/repository/{ownerUsername}/{repositoryName}/content/readme`** — Get readme file content public

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `repoType` | query | `string` | No | allow: dataset, model |

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
resp, err := client.Publics.Public.GetPublicRepositoryByOwnerusernameByRepositorynameContentReadme(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicSearch`

**`GET /public/search`** — Multiple Search

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
resp, err := client.Publics.Public.GetPublicSearch(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicSearchUser`

**`GET /public/search/user`** — Search Public Users

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `search` | query | `string` | Yes |  |

**Responses**

**200 OK** — `response.UserListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.UserListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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
resp, err := client.Publics.Public.GetPublicSearchUser(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicTokenPrice`

**`GET /public/token/price`** — Get Aioz Price

**Responses**

**200 OK** — `response.GetTokenPriceResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `utils.GetTokenInfoData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`utils.GetTokenInfoData`**

| Field | Type | Description |
| --- | --- | --- |
| `change_percentage_24h` | `number` |  |
| `circulating_supply` | `number` |  |
| `current_price` | `number` |  |
| `fully_diluted_valuation` | `number` |  |
| `high_24_h` | `number` |  |
| `last_updated` | `string` |  |
| `low_24_h` | `number` |  |
| `market_cap` | `number` |  |
| `rank` | `integer` |  |
| `total_supply` | `number` |  |
| `volume_24_h` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicTokenPrice(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicUserByUsername`

**`GET /public/user/{username}`** — Get user's info

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | username |

**Responses**

**200 OK** — `response.GetUserByUsernameAndGuestIdResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetUserByUsernameAndGuestIdData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetUserByUsernameAndGuestIdData`**

| Field | Type | Description |
| --- | --- | --- |
| `user` | `models.LiteUser` |  |

**`models.LiteUser`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_request_to_join` | `boolean` |  |
| `avatar_url` | `string` |  |
| `bio` | `string` |  |
| `blocked` | `boolean` |  |
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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicUserByUsername(ctx, "<username>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicUserByUsernameExisted`

**`GET /public/user/{username}/existed`** — Check if a username have already existed

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | username |

**Responses**

**200 OK** — `response.CheckUsernameExistResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckUsernameExistData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckUsernameExistData`**

| Field | Type | Description |
| --- | --- | --- |
| `existed` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicUserByUsernameExisted(ctx, "<username>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicUserByUsernameMedals`

**`GET /public/user/{username}/medals`** — Get user medals by medal name

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Username |
| `limit` | query | `integer` | No |  |
| `medalName` | query | `string` | Yes |  |
| `offset` | query | `integer` | No |  |

**Responses**

**200 OK** — `response.UserMedalListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.UserMedalListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.UserMedalListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.UserMedal]` |  |
| `total` | `integer` |  |

**`models.UserMedal`**

| Field | Type | Description |
| --- | --- | --- |
| `competition` | `models.Competition` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `medal_name` | `string` |  |
| `rank` | `integer` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Medalss.Medals.GetPublicUserByUsernameMedals(ctx, "<username>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicUserByUsernameMedalsStatistics`

**`GET /public/user/{username}/medals/statistics`** — Get user medal statistics by username

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | Username |

**Responses**

**200 OK** — `response.MedalStatisticResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.MedalStatistic` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.MedalStatistic`**

| Field | Type | Description |
| --- | --- | --- |
| `count` | `integer` |  |
| `medal_name` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Medalss.Medals.GetPublicUserByUsernameMedalsStatistics(ctx, "<username>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicUserByUsernameOrganizations`

**`GET /public/user/{username}/organizations`** — Get public user's organizations by username

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | path | `string` | Yes | username |
| `keyword` | query | `string` | No |  |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |

**Responses**

**200 OK** — `response.GetUserOrganizationsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetUserOrganizationsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetUserOrganizationsData`**

| Field | Type | Description |
| --- | --- | --- |
| `organizations` | `array[models.OrganizationInfo]` |  |
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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Publics.Public.GetPublicUserByUsernameOrganizations(ctx, "<username>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
