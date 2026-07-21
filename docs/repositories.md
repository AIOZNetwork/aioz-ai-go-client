# Repositories — Go SDK

> Repository management — code and file repositories. Requires: `x-api-key` header.

Reference: [SDK Usage Guide](../README.md#sdk-usage-guide) | [Package README](../README.md)

---

### `GetRepositoryByOwnerusernameByRepositorynameContentReadme`

**`GET /api-key/repository/{ownerUsername}/{repositoryName}/content/readme`** — Get readme file content

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

**Error Responses**

| Status | Description |
| --- | --- |
| 400 | Bad Request — [response.FailResponse](../README.md#common-response-types) |
| 500 | Internal Server Error — [response.ErrorResponse](../README.md#common-response-types) |

**Example**

```go
ctx := context.Background()
resp, err := client.Repositories().Repository.GetRepositoryByOwnerusernameByRepositorynameContentReadme(repository.NewGetRepositoryByOwnerusernameByRepositorynameContentReadmeParams().WithContext(ctx).WithOwnerUsername("<ownerUsername>").WithRepositoryName("<repositoryName>"), nil, io.Discard)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("%+v\n", resp)
```

---
