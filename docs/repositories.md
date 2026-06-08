# Repositories — Go SDK

> Repository management — code and file repositories. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetRepositoryModelCommitByCommitShaHistory`

**`GET /api-key/repository/model/commit/{commit_sha}/history`** — Get Checked Commit By Id (model) by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `commit_sha` | path | `string` | Yes | Commit's SHA |

**Responses**

**200 OK** — `response.CheckedCommitListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckedCommitListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckedCommitListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.CheckedCommit]` |  |
| `total` | `integer` |  |

**`models.CheckedCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `commit_hash` | `string` |  |
| `commit_message` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `model_id` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `valid` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryModelCommitByCommitShaHistory(ctx, "<commit_sha>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryModelByIDCommitHistory`

**`GET /api-key/repository/model/{id}/commit/history`** — Get List Checked Commit (model) by api-key

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

**200 OK** — `response.CheckedCommitListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.CheckedCommitListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.CheckedCommitListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.CheckedCommit]` |  |
| `total` | `integer` |  |

**`models.CheckedCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `commit_hash` | `string` |  |
| `commit_message` | `string` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `model_id` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `username` | `string` |  |
| `valid` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryModelByIDCommitHistory(ctx, "<id>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositoryname`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}`** — Get repository by name by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `repoType` | query | `string` | No | allow: dataset, model |

**Responses**

**200 OK** — `response.GetRepositoryResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetRepositoryData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetRepositoryData`**

| Field | Type | Description |
| --- | --- | --- |
| `repository` | `models.Repository` |  |

**`models.Repository`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_merge` | `boolean` |  |
| `allow_rebase` | `boolean` |  |
| `allow_rebase_merge` | `boolean` |  |
| `allow_squash` | `boolean` |  |
| `archived` | `boolean` |  |
| `clone_url` | `string` |  |
| `created` | `string` |  |
| `default_branch` | `string` |  |
| `default_merge_style` | `string` |  |
| `description` | `string` |  |
| `empty` | `boolean` |  |
| `fork` | `boolean` |  |
| `forks` | `integer` |  |
| `full_name` | `string` |  |
| `has_pull_requests` | `boolean` |  |
| `id` | `integer` |  |
| `mirror` | `boolean` |  |
| `name` | `string` |  |
| `original_url` | `string` |  |
| `owner` | `models.User` |  |
| `parent` | `models.Repository` |  |
| `permissions` | `models.Permission` |  |
| `private` | `boolean` |  |
| `size` | `integer` |  |
| `ssh_url` | `string` |  |
| `stars` | `integer` |  |
| `template` | `boolean` |  |
| `updated` | `string` |  |
| `watchers` | `integer` |  |
| `website` | `string` |  |

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

**`models.Permission`**

| Field | Type | Description |
| --- | --- | --- |
| `admin` | `boolean` |  |
| `pull` | `boolean` |  |
| `push` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositoryname(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameBranches`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/branches`** — Get repository branches by repository name by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |
| `repoType` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetBranchesResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetBranchesData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetBranchesData`**

| Field | Type | Description |
| --- | --- | --- |
| `branches` | `array[models.Branch]` |  |

**`models.Branch`**

| Field | Type | Description |
| --- | --- | --- |
| `commit` | `models.PayloadCommit` |  |
| `effective_branch_protection_name` | `string` |  |
| `enable_status_check` | `boolean` |  |
| `name` | `string` |  |
| `protected` | `boolean` |  |
| `required_approvals` | `integer` |  |
| `status_check_contexts` | `array[string]` |  |
| `user_can_merge` | `boolean` |  |
| `user_can_push` | `boolean` |  |

**`models.PayloadCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `added` | `array[string]` |  |
| `author` | `models.PayloadUser` |  |
| `committer` | `models.PayloadUser` |  |
| `id` | `string` |  |
| `message` | `string` |  |
| `modified` | `array[string]` |  |
| `removed` | `array[string]` |  |
| `timestamp` | `string` |  |

**`models.PayloadUser`**

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

**`models.PayloadUser`**

| Field | Type | Description |
| --- | --- | --- |
| `email` | `string` |  |
| `name` | `string` |  |
| `username` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameBranches(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostRepositoryByOwnerusernameByRepositorynameBuy`

**`POST /api-key/repository/{ownerUsername}/{repositoryName}/buy`** — Unlock Repository By User by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |

**Request Body** — `request.DownloadRepoZipRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `ref` | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | `string` | No |  |
| `zipType` | `string` | No | ZipType is the type of the zip file (optional) |

**Responses**

**200 OK** — `response.UnlockRepoByUserResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetCommitDiffData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetCommitDiffData`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.DownloadRepoZipRequest{
    Ref: "...",  // string
    Repotype: "...",  // string
    Ziptype: "...",  // string
}
resp, err := client.Repositorys.Repository.PostRepositoryByOwnerusernameByRepositorynameBuy(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameCollaborators`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/collaborators`** — Get repository collaborators by repository name by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |
| `repoType` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetCollaboratorsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetCollaboratorsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetCollaboratorsData`**

| Field | Type | Description |
| --- | --- | --- |
| `collaborators` | `array[models.User]` |  |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameCollaborators(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameCommit`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/commit`** — Get repository single commit by repository name and commit sha by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `repoType` | query | `string` | No |  |
| `sha` | query | `string` | Yes |  |

**Responses**

**200 OK** — `response.GetCommitHistoryResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetCommitHistoryData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetCommitHistoryData`**

| Field | Type | Description |
| --- | --- | --- |
| `commit` | `array[models.Commit]` |  |
| `has_more` | `boolean` |  |
| `last_page` | `integer` |  |

**`models.Commit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.User` |  |
| `commit_affected_files` | `array[string]` |  |
| `commit_data` | `models.User` |  |
| `commit_meta` | `models.CommitMeta` |  |
| `parents` | `array[models.CommitMeta]` |  |
| `repo_commit` | `models.RepoCommit` |  |
| `stats` | `models.CommitStats` |  |

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

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.RepoCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.CommitUser` |  |
| `committer` | `models.CommitUser` |  |
| `message` | `string` |  |
| `tree` | `models.CommitMeta` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitStats`**

| Field | Type | Description |
| --- | --- | --- |
| `additions` | `integer` |  |
| `deletions` | `integer` |  |
| `total` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameCommit(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameCommitHistory`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/commit/history`** — Get commit history by repository name and branch name

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |
| `path` | query | `string` | No | Path is the path of the file (optional) |
| `repoType` | query | `string` | No |  |
| `sha` | query | `string` | No | Sha is the sha of the commit (optional) |

**Responses**

**200 OK** — `response.GetCommitHistoryResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetCommitHistoryData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetCommitHistoryData`**

| Field | Type | Description |
| --- | --- | --- |
| `commit` | `array[models.Commit]` |  |
| `has_more` | `boolean` |  |
| `last_page` | `integer` |  |

**`models.Commit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.User` |  |
| `commit_affected_files` | `array[string]` |  |
| `commit_data` | `models.User` |  |
| `commit_meta` | `models.CommitMeta` |  |
| `parents` | `array[models.CommitMeta]` |  |
| `repo_commit` | `models.RepoCommit` |  |
| `stats` | `models.CommitStats` |  |

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

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.RepoCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.CommitUser` |  |
| `committer` | `models.CommitUser` |  |
| `message` | `string` |  |
| `tree` | `models.CommitMeta` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitStats`**

| Field | Type | Description |
| --- | --- | --- |
| `additions` | `integer` |  |
| `deletions` | `integer` |  |
| `total` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameCommitHistory(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameCommitTree`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/commit/tree`** — Get commit tree by repository name and commit sha by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `ref` | query | `string` | Yes | Ref is the ref of the commit (required) |
| `repoType` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetCommitTreeResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetCommitTreeData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetCommitTreeData`**

| Field | Type | Description |
| --- | --- | --- |
| `tree` | `models.CommitTree` |  |

**`models.CommitTree`**

| Field | Type | Description |
| --- | --- | --- |
| `page` | `integer` |  |
| `sha` | `string` |  |
| `total_count` | `integer` |  |
| `tree` | `array[models.GitEntry]` |  |
| `truncated` | `boolean` |  |
| `url` | `string` |  |

**`models.GitEntry`**

| Field | Type | Description |
| --- | --- | --- |
| `mode` | `string` |  |
| `path` | `string` |  |
| `sha` | `string` |  |
| `size` | `integer` |  |
| `type` | `string` |  |
| `url` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameCommitTree(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameContent`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content`** — Get file content by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `path` | query | `string` | No | Path is the path of the file (required) |
| `ref` | query | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | query | `string` | No |  |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameContent(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PutRepositoryByOwnerusernameByRepositorynameContent`

**`PUT /api-key/repository/{ownerUsername}/{repositoryName}/content`** — Modify repository content by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |

**Request Body** — `request.CreateRepositoryContentRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `branch_name` | `string` | No | BranchName is the name of the branch (required) |
| `email` | `string` | No | Email is the email of the committer (optional) (default: author's email) |
| `files` | `array[models.File]` | No | Files is the list of files to be created (required) (min:1, max:5, max file size: 50MB) |
| `message` | `string` | No | Commit message is the commit message (required) |
| `new_branch_name` | `string` | No | NewBranchName is the name of the new branch (default usually master) (optional) |
| `repo_type` | `string` | No |  |
| `signoff` | `boolean` | No | Signoff is the flag to indicate if the commit is signed off (default: false) (optional) |
| `timestamp` | `string` | No | Timestamp is the timestamp of the commit (default: now) (optional) |

**`models.File`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` | Content is the content of the file in base64 format (optional) |
| `operation` | `string` | Operation is the operation to be performed on the file (optional) |
| `path` | `string` | Path is the path of the file (optional) |
| `sha` | `string` | Sha is the sha of the file (optional) |

**Responses**

**200 OK** — `response.ModifyRepoContentResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.ModifyRepoContentData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.ModifyRepoContentData`**

| Field | Type | Description |
| --- | --- | --- |
| `commit` | `models.Commit` |  |

**`models.Commit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.User` |  |
| `commit_affected_files` | `array[string]` |  |
| `commit_data` | `models.User` |  |
| `commit_meta` | `models.CommitMeta` |  |
| `parents` | `array[models.CommitMeta]` |  |
| `repo_commit` | `models.RepoCommit` |  |
| `stats` | `models.CommitStats` |  |

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

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.RepoCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `author` | `models.CommitUser` |  |
| `committer` | `models.CommitUser` |  |
| `message` | `string` |  |
| `tree` | `models.CommitMeta` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitStats`**

| Field | Type | Description |
| --- | --- | --- |
| `additions` | `integer` |  |
| `deletions` | `integer` |  |
| `total` | `integer` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.CreateRepositoryContentRequest{
    BranchName: "...",  // string
    Email: "...",  // string
    Files: "...",  // array[models.File]
    Message: "...",  // string
    NewBranchName: "...",  // string
}
resp, err := client.Repositorys.Repository.PutRepositoryByOwnerusernameByRepositorynameContent(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteRepositoryByOwnerusernameByRepositorynameContent`

**`DELETE /api-key/repository/{ownerUsername}/{repositoryName}/content`** — Delete single repository content by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |

**Request Body** — `request.DeleteRepositoryContentRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `branch_name` | `string` | Yes | BranchName is the name of the branch (required) |
| `message` | `string` | Yes | Commit message is the commit message (required) |
| `new_branch_name` | `string` | No | NewBranchName is the name of the new branch (default usually master) (optional) |
| `path` | `string` | No | Path is the path of the file to be deleted (optional) |
| `sha` | `string` | Yes | Sha is the sha of the file to be deleted (required) |
| `signoff` | `boolean` | No | Signoff is the flag to indicate if the commit is signed off (default: false) (optional) |

**Responses**

**200 OK** — `response.DeleteRepositoryContentResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.DeleteRepositoryContentData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.DeleteRepositoryContentData`**

| Field | Type | Description |
| --- | --- | --- |
| `commit` | `models.FileCommit` |  |

**`models.FileCommit`**

| Field | Type | Description |
| --- | --- | --- |
| `commit_data` | `models.CommitUser` |  |
| `commit_meta` | `models.CommitMeta` |  |
| `commit_user` | `models.CommitUser` |  |
| `message` | `string` |  |
| `parents` | `array[models.CommitMeta]` |  |
| `tree` | `models.CommitMeta` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitUser`**

| Field | Type | Description |
| --- | --- | --- |
| `date` | `string` |  |
| `email` | `string` |  |
| `name` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**`models.CommitMeta`**

| Field | Type | Description |
| --- | --- | --- |
| `created` | `string` |  |
| `sha` | `string` |  |
| `url` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.DeleteRepositoryContentRequest{
    BranchName: "...",  // string  // required
    Message: "...",  // string  // required
    NewBranchName: "...",  // string
    Path: "...",  // string
    Sha: "...",  // string  // required
}
resp, err := client.Repositorys.Repository.DeleteRepositoryByOwnerusernameByRepositorynameContent(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameContentMetadata`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content/metadata`** — Get repository content metadata by repository name, branch name and path by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `path` | query | `string` | No |  |
| `ref` | query | `string` | No |  |
| `repoType` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetRepositoryContentMetadataResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetRepositoryContentMetadataData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetRepositoryContentMetadataData`**

| Field | Type | Description |
| --- | --- | --- |
| `contents` | `array[models.Content]` |  |
| `total` | `integer` |  |

**`models.Content`**

| Field | Type | Description |
| --- | --- | --- |
| `content` | `string` |  |
| `download_url` | `string` |  |
| `encoding` | `string` |  |
| `git_url` | `string` |  |
| `is_lfs` | `boolean` |  |
| `last_commit_at` | `string` |  |
| `last_commit_message` | `string` |  |
| `last_commit_sha` | `string` |  |
| `links` | `models.FileLink` |  |
| `name` | `string` |  |
| `path` | `string` |  |
| `sha` | `string` |  |
| `size` | `integer` |  |
| `submodule_git_url` | `string` |  |
| `target` | `string` |  |
| `type` | `string` |  |
| `url` | `string` |  |

**`models.FileLink`**

| Field | Type | Description |
| --- | --- | --- |
| `git` | `string` |  |
| `html` | `string` |  |
| `self` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameContentMetadata(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameContentPlayground`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content/playground`** — Get playground file content (model only) by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameContentPlayground(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameContentPreDownload`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content/pre-download`** — Calculate Cost To Download Repo Content by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `path` | query | `string` | No | Path is the path of the file (required) |
| `ref` | query | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | query | `string` | No |  |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameContentPreDownload(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameContentReadme`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content/readme`** — Get readme file content by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `path` | query | `string` | No | Path is the path of the file (required) |
| `ref` | query | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | query | `string` | No |  |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameContentReadme(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynameDownload`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/download`** — Download repository zip by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `ref` | query | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | query | `string` | No |  |
| `zipType` | query | `string` | No | ZipType is the type of the zip file (optional) |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynameDownload(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetRepositoryByOwnerusernameByRepositorynamePreDownload`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/pre-download`** — Check Available Free Download Repo Zip by api-key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `ownerUsername` | path | `string` | Yes | repository's owner |
| `repositoryName` | path | `string` | Yes | repository's name |
| `ref` | query | `string` | No | Ref is the ref of the commit, name of the branch, default: main (optional) |
| `repoType` | query | `string` | No |  |
| `zipType` | query | `string` | No | ZipType is the type of the zip file (optional) |

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
resp, err := client.Repositorys.Repository.GetRepositoryByOwnerusernameByRepositorynamePreDownload(ctx, "<ownerUsername>", "<repositoryName>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
