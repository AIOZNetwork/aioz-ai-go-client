# Organizations — Go SDK

> Organization management — teams, members, and wallets. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `PostAPIKeyOrganization`

**`POST /api-key/organization`** — Create new organization By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Request Body** — `request.CreateOrganizationRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar` | `string` | No | Avatar is the avatar of the organization (optional) |
| `bio` | `string` | No | Bio is the full name of the organization (optional) |
| `full_name` | `string` | No | FullName is the full name of the organization (optional) |
| `github_name` | `string` | No | GithubName is the github name of the organization (optional) |
| `home_page` | `string` | No | HomePage is the home page of the organization (optional) |
| `interests` | `string` | No | Interests is the list of interests of the organization (optional) |
| `twitter_name` | `string` | No | TwitterName is the twitter name of the organization (optional) |
| `type` | `string` | Yes | Type is the type of the organization (allow: company, university, classroom, non_profit, community) (required). One of: `company`, `university`, `classroom`, `non_profit`, `community` |
| `username` | `string` | Yes | Username is the username of the organization (required) |
| `visibility` | `string` | No | One of: `public`, `private` |

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
req := &models.CreateOrganizationRequest{
    Avatar: "...",  // string
    Bio: "...",  // string
    FullName: "...",  // string
    GithubName: "...",  // string
    HomePage: "...",  // string
}
resp, err := client.Organizations.Organization.PostAPIKeyOrganization(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationList`

**`GET /api-key/organization/list`** — Get organizations By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `keyword` | query | `string` | No |  |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `type` | query | `string` | No |  |

**Responses**

**200 OK** — `response.OrganizationListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.OrganizationListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationList(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrg`

**`GET /api-key/organization/{org}`** — Get organization By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Responses**

**200 OK** — `response.OrganizationResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.OrganizationInfo` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrg(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyOrganizationOrg`

**`DELETE /api-key/organization/{org}`** — Delete organization By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

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
resp, err := client.Organizations.Organization.DeleteAPIKeyOrganizationOrg(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PatchAPIKeyOrganizationOrgInfo`

**`PATCH /api-key/organization/{org}/info`** — Update organization's info By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.UpdateOrganizationInfoRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `avatar` | `string` | No | Avatar is the avatar of the organization (optional) |
| `bio` | `string` | No | Bio is the full name of the organization (optional) |
| `full_name` | `string` | No | FullName is the full name of the organization (optional) |
| `github_link` | `string` | No | GithubLink is the github link of the organization (optional) |
| `github_name` | `string` | No | GithubName is the github name of the organization (optional) |
| `home_page` | `string` | No | HomePage is the home page of the organization (optional) |
| `interests` | `string` | No | Interests is the list of interests of the organization (optional) |
| `twitter_link` | `string` | No | TwitterLink is the twitter link of the organization (optional) |
| `twitter_name` | `string` | No | TwitterName is the twitter name of the organization (optional) |
| `type` | `string` | No | Type is the type of the organization (optional) |
| `visibility` | `string` | Yes | One of: `public`, `private` |

**Responses**

**200 OK** — `response.OrganizationResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.OrganizationInfo` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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
req := &models.UpdateOrganizationInfoRequest{
    Avatar: "...",  // string
    Bio: "...",  // string
    FullName: "...",  // string
    GithubLink: "...",  // string
    GithubName: "...",  // string
}
resp, err := client.Organizations.Organization.PatchAPIKeyOrganizationOrgInfo(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgIsMember`

**`GET /api-key/organization/{org}/is-member`** — Check if user is member of organization By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `username` | query | `string` | Yes | data |
| `org` | path | `string` | Yes | organization's username |

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
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrgIsMember(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyOrganizationOrgLeave`

**`DELETE /api-key/organization/{org}/leave`** — Leave Organization By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

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
resp, err := client.Organizations.Organization.DeleteAPIKeyOrganizationOrgLeave(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgMembers`

**`GET /api-key/organization/{org}/members`** — Get organization's members By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |

**Responses**

**200 OK** — `response.GetPublicOrganizationMembersResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetPublicOrganizationMembersData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetPublicOrganizationMembersData`**

| Field | Type | Description |
| --- | --- | --- |
| `members` | `array[models.LiteUser]` |  |
| `total` | `integer` |  |

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
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrgMembers(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgOffers`

**`GET /api-key/organization/{org}/offers`** — Get organization's offers By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |
| `keyword` | query | `string` | No |  |
| `offerType` | query | `string` | Yes |  |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |

**Responses**

**200 OK** — `response.GetOffersResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetOffersData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetOffersData`**

| Field | Type | Description |
| --- | --- | --- |
| `offers` | `array[models.Offer]` |  |

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
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrgOffers(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgPermission`

**`GET /api-key/organization/{org}/permission`** — Get user organization permission By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetUserOrganizationPermissionRequest`

_No fields defined._

**Responses**

**200 OK** — `response.GetUserOrganizationPermissionResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetUserOrganizationPermissionData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetUserOrganizationPermissionData`**

| Field | Type | Description |
| --- | --- | --- |
| `permission` | `models.OrgPermission` |  |

**`models.OrgPermission`**

| Field | Type | Description |
| --- | --- | --- |
| `can_create` | `boolean` |  |
| `can_delete` | `boolean` |  |
| `can_edit` | `boolean` |  |
| `can_read` | `boolean` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetUserOrganizationPermissionRequest{
    
}
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrgPermission(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgSetting`

**`GET /api-key/organization/{org}/setting`** — Get organization's setting By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Responses**

**200 OK** — `response.GetOrganizationSettingResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetOrganizationSettingData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetOrganizationSettingData`**

| Field | Type | Description |
| --- | --- | --- |
| `allow_join_by_link` | `boolean` |  |
| `allow_request_to_join` | `boolean` |  |
| `auto_approve_join_request` | `boolean` |  |
| `default_role` | `string` |  |
| `domain` | `string` |  |
| `email` | `string` |  |
| `join_id` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Organizations.Organization.GetAPIKeyOrganizationOrgSetting(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PatchAPIKeyOrganizationOrgSetting`

**`PATCH /api-key/organization/{org}/setting`** — Update organization's join settings By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.UpdateJoinSettingsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `allow_join_by_link` | `boolean` | Yes | AllowJoinByLink is the flag to indicate if the organization allows join by link (required) |
| `allow_request_to_join` | `boolean` | Yes | AllowRequestToJoin is the flag to indicate if the organization allows request to join (required) |
| `auto_approve_join_request` | `boolean` | Yes | AutoApproveJoinRequest is the flag to indicate if the organization auto approves join request (required) |
| `default_role` | `string` | Yes | DefaultRole is the default role when approve new member to organization (required) |

**Responses**

**200 OK** — `response.OrganizationResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.OrganizationInfo` |  |
| `message` | `string` |  |
| `status` | `string` |  |

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
req := &models.UpdateJoinSettingsRequest{
    AllowJoinByLink: "...",  // boolean  // required
    AllowRequestToJoin: "...",  // boolean  // required
    AutoApproveJoinRequest: "...",  // boolean  // required
    DefaultRole: "...",  // string  // required
}
resp, err := client.Organizations.Organization.PatchAPIKeyOrganizationOrgSetting(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOrganizationOrgStatisticsEarnings`

**`POST /api-key/organization/{org}/statistics/earnings`** — Get Organzition Earnings Statistics By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetTransactionAnalyticsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `to` | `string` | No | 2024-05-07 15:04:05 |

**Responses**

**200 OK** — `response.UserEarningsStatisticsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.UserEarningsStatistics` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.UserEarningsStatistics`**

| Field | Type | Description |
| --- | --- | --- |
| `api_calls` | `integer` |  |
| `api_calls_amount` | `number` |  |
| `dataset_used` | `integer` |  |
| `dataset_used_amount` | `number` |  |
| `model_used` | `integer` |  |
| `model_used_amount` | `number` |  |
| `playground_used` | `integer` |  |
| `playground_used_amount` | `number` |  |
| `total` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetTransactionAnalyticsRequest{
    From: "...",  // string
    To: "...",  // string
}
resp, err := client.Organizations.Wallet.PostAPIKeyOrganizationOrgStatisticsEarnings(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOrganizationOrgStatisticsSpendingCost`

**`POST /api-key/organization/{org}/statistics/spending-cost`** — Get Organzition Spending Cost Statistics By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetOrganizationSpendingCostStatisticsRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `to` | `string` | No | 2023-05-07 15:04:05 |

**Responses**

**200 OK** — `response.UserSpendingCostStatisticsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.UserSpendingCostStatistics` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.UserSpendingCostStatistics`**

| Field | Type | Description |
| --- | --- | --- |
| `api_calls` | `integer` |  |
| `api_calls_amount` | `number` |  |
| `bandwidth` | `number` |  |
| `bandwidth_amount` | `number` |  |
| `dataset_used` | `integer` |  |
| `dataset_used_amount` | `number` |  |
| `model_used` | `integer` |  |
| `model_used_amount` | `number` |  |
| `playground_used` | `integer` |  |
| `playground_used_amount` | `number` |  |
| `storage` | `number` |  |
| `storage_amount` | `number` |  |
| `total` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetOrganizationSpendingCostStatisticsRequest{
    From: "...",  // string
    To: "...",  // string
}
resp, err := client.Organizations.Wallet.PostAPIKeyOrganizationOrgStatisticsSpendingCost(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgWalletDepositHistory`

**`GET /api-key/organization/{org}/wallet/deposit/history`** — Get List Org Deposit History By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetListOrgDepositHistoryByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |

**Responses**

**200 OK** — `response.TxInListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.TxInListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.TxInListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.TxIn]` |  |
| `total` | `integer` |  |

**`models.TxIn`**

| Field | Type | Description |
| --- | --- | --- |
| `aioz_usd_rate` | `number` |  |
| `amount` | `string` |  |
| `block_number` | `integer` |  |
| `cosmos_tx_hash` | `string` |  |
| `created_at` | `string` |  |
| `credit` | `string` |  |
| `evm_tx_hash` | `string` |  |
| `sender` | `string` |  |
| `status` | `boolean` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetListOrgDepositHistoryByOwnerRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Organizations.Wallet.GetAPIKeyOrganizationOrgWalletDepositHistory(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOrganizationOrgWalletTransactionAnalytics`

**`POST /api-key/organization/{org}/wallet/transaction/analytics`** — Get Org Transaction Analytics By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetOrgTransactionAnalyticsByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `to` | `string` | No | 2023-05-07 15:04:05 |

**Responses**

**200 OK** — `response.TransactionAnalyticsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.TransactionAnalytics` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.TransactionAnalytics`**

| Field | Type | Description |
| --- | --- | --- |
| `earnings_analytics` | `models.EarningsAnalytics` |  |
| `spending_analytics` | `models.SpendingAnalytics` |  |

**`models.EarningsAnalytics`**

| Field | Type | Description |
| --- | --- | --- |
| `api_call_earnings` | `number` |  |
| `dataset_earnings` | `number` |  |
| `model_earnings` | `number` |  |
| `playground_earnings` | `number` |  |
| `total_earnings` | `number` |  |

**`models.SpendingAnalytics`**

| Field | Type | Description |
| --- | --- | --- |
| `api_call_used` | `number` |  |
| `bandwidth` | `number` |  |
| `dataset_used` | `number` |  |
| `model_used` | `number` |  |
| `package` | `number` |  |
| `playground_used` | `number` |  |
| `storage` | `number` |  |
| `total_spending` | `number` |  |
| `verify_model` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetOrgTransactionAnalyticsByOwnerRequest{
    From: "...",  // string
    To: "...",  // string
}
resp, err := client.Organizations.Wallet.PostAPIKeyOrganizationOrgWalletTransactionAnalytics(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgWalletTransactionHistory`

**`GET /api-key/organization/{org}/wallet/transaction/history`** — Get List Org Transaction History By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetListOrgTransactionHistoryByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `type` | `string` | No |  |

**Responses**

**200 OK** — `response.TransactionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.TransactionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.TransactionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Transaction]` |  |
| `total` | `integer` |  |
| `total_amount` | `number` |  |
| `total_value` | `number` |  |

**`models.Transaction`**

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` |  |
| `aioz_usd_rate` | `number` |  |
| `amount` | `number` |  |
| `burn` | `number` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `object_id` | `string` |  |
| `object_type` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `value` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetListOrgTransactionHistoryByOwnerRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
    Type: "...",  // string
}
resp, err := client.Organizations.Wallet.GetAPIKeyOrganizationOrgWalletTransactionHistory(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOrganizationOrgWalletTransactionRecent`

**`POST /api-key/organization/{org}/wallet/transaction/recent`** — Get List Org Recent Transaction By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetListOrgRecentTransactionByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `from` | `string` | No | 2023-05-07 15:04:05 |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |
| `to` | `string` | No | 2023-05-07 15:04:05 |
| `type` | `string` | No |  |

**Responses**

**200 OK** — `response.TransactionListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.TransactionListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.TransactionListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.Transaction]` |  |
| `total` | `integer` |  |
| `total_amount` | `number` |  |
| `total_value` | `number` |  |

**`models.Transaction`**

| Field | Type | Description |
| --- | --- | --- |
| `action` | `string` |  |
| `aioz_usd_rate` | `number` |  |
| `amount` | `number` |  |
| `burn` | `number` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `object_id` | `string` |  |
| `object_type` | `string` |  |
| `updated_at` | `string` |  |
| `user_id` | `string` |  |
| `value` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetListOrgRecentTransactionByOwnerRequest{
    From: "...",  // string
    Limit: "...",  // integer
    Offset: "...",  // integer
    To: "...",  // string
    Type: "...",  // string
}
resp, err := client.Organizations.Wallet.PostAPIKeyOrganizationOrgWalletTransactionRecent(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PostAPIKeyOrganizationOrgWalletWithdrawEarnings`

**`POST /api-key/organization/{org}/wallet/withdraw/earnings`** — Withdraw Org Earnings By Owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.WithdrawOrgEarningsByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `amount` | `string` | Yes |  |
| `to_wallet` | `string` | Yes |  |

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
req := &models.WithdrawOrgEarningsByOwnerRequest{
    Amount: "...",  // string  // required
    ToWallet: "...",  // string  // required
}
resp, err := client.Organizations.Wallet.PostAPIKeyOrganizationOrgWalletWithdrawEarnings(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetAPIKeyOrganizationOrgWalletWithdrawHistory`

**`GET /api-key/organization/{org}/wallet/withdraw/history`** — Get list org withdrawal history by owner By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.GetListOrgWithdrawalHistoryByOwnerRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `limit` | `integer` | No |  |
| `offset` | `integer` | No |  |

**Responses**

**200 OK** — `response.WithdrawalListResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.WithdrawalListData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.WithdrawalListData`**

| Field | Type | Description |
| --- | --- | --- |
| `records` | `array[models.TreasuryWithdrawal]` |  |
| `total` | `integer` |  |

**`models.TreasuryWithdrawal`**

| Field | Type | Description |
| --- | --- | --- |
| `amount` | `number` |  |
| `created_at` | `string` |  |
| `evm_tx_hash` | `string` |  |
| `from_service` | `string` |  |
| `id` | `string` |  |
| `recipient` | `string` |  |
| `sender` | `string` |  |
| `updated_at` | `string` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
req := &models.GetListOrgWithdrawalHistoryByOwnerRequest{
    Limit: "...",  // integer
    Offset: "...",  // integer
}
resp, err := client.Organizations.Wallet.GetAPIKeyOrganizationOrgWalletWithdrawHistory(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `DeleteAPIKeyOrganizationOrgMember`

**`DELETE /api-key/organization/{org}/{member}`** — Delete organization's member By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |
| `member` | path | `string` | Yes | member's username |

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
resp, err := client.Organizations.Organization.DeleteAPIKeyOrganizationOrgMember(ctx, "<org>", "<member>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `PatchAPIKeyOrganizationOrgMember`

**`PATCH /api-key/organization/{org}/{member}`** — Change member's role By Api Key

**Headers**

| Header | Value | Required |
| --- | --- | --- |
| `x-api-key` | Your API key | Yes |

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Request Body** — `request.ChangeMemberRoleRequest`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `role` | `string` | Yes |  |

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
req := &models.ChangeMemberRoleRequest{
    Role: "...",  // string  // required
}
resp, err := client.Organizations.Organization.PatchAPIKeyOrganizationOrgMember(ctx, req)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicOrganization`

**`GET /public/organization`** — Get public organizations

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `limit` | query | `integer` | No |  |
| `offset` | query | `integer` | No |  |
| `sort` | query | `string` | No |  |

**Responses**

**200 OK** — `response.GetLiteOrganizationsResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetLiteOrganizationsData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetLiteOrganizationsData`**

| Field | Type | Description |
| --- | --- | --- |
| `organizations` | `array[models.LiteOrganization]` |  |
| `total` | `integer` |  |

**`models.LiteOrganization`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `bio` | `string` |  |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `datasets_count` | `integer` |  |
| `full_name` | `string` |  |
| `github_link` | `string` |  |
| `github_name` | `string` |  |
| `home_page` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `join_id` | `string` |  |
| `models_count` | `integer` |  |
| `org_team` | `string` |  |
| `total_repos_size` | `integer` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |
| `updated_by` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |
| `wallet` | `models.Wallet` |  |
| `wallet_address` | `string` |  |

**`models.Wallet`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `number` |  |
| `debt` | `number` |  |
| `earnings` | `number` |  |
| `free_balance` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Organizations.Organization.GetPublicOrganization(ctx)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicOrganizationOrg`

**`GET /public/organization/{org}`** — Get organization detail

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |

**Responses**

**200 OK** — `response.LiteOrganizationResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `models.LiteOrganization` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`models.LiteOrganization`**

| Field | Type | Description |
| --- | --- | --- |
| `avatar` | `string` |  |
| `bio` | `string` |  |
| `created_at` | `string` |  |
| `created_by` | `string` |  |
| `datasets_count` | `integer` |  |
| `full_name` | `string` |  |
| `github_link` | `string` |  |
| `github_name` | `string` |  |
| `home_page` | `string` |  |
| `id` | `string` |  |
| `interests` | `string` |  |
| `join_id` | `string` |  |
| `models_count` | `integer` |  |
| `org_team` | `string` |  |
| `total_repos_size` | `integer` |  |
| `twitter_link` | `string` |  |
| `twitter_name` | `string` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |
| `updated_by` | `string` |  |
| `username` | `string` |  |
| `verified` | `boolean` |  |
| `visibility` | `string` |  |
| `wallet` | `models.Wallet` |  |
| `wallet_address` | `string` |  |

**`models.Wallet`**

| Field | Type | Description |
| --- | --- | --- |
| `balance` | `number` |  |
| `debt` | `number` |  |
| `earnings` | `number` |  |
| `free_balance` | `number` |  |

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Organizations.Organization.GetPublicOrganizationOrg(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---

### `GetPublicOrganizationOrgMembers`

**`GET /public/organization/{org}/members`** — Get pulic organization's members

**Parameters**

| Name | Location | Type | Required | Description |
| --- | --- | --- | --- | --- |
| `org` | path | `string` | Yes | organization's username |
| `page` | query | `integer` | No | Page is the page number (default: 1) (optional) |
| `pageSize` | query | `integer` | No | PageSize is the page size (default: 10) (optional) |

**Responses**

**200 OK** — `response.GetPublicOrganizationMembersResponse`

| Field | Type | Description |
| --- | --- | --- |
| `data` | `response.GetPublicOrganizationMembersData` |  |
| `message` | `string` |  |
| `status` | `string` |  |

**`response.GetPublicOrganizationMembersData`**

| Field | Type | Description |
| --- | --- | --- |
| `members` | `array[models.LiteUser]` |  |
| `total` | `integer` |  |

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
resp, err := client.Organizations.Organization.GetPublicOrganizationOrgMembers(ctx, "<org>")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
