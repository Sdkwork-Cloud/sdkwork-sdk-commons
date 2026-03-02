# sdk-common-go

Common foundation package for generated Go SDKs.

## Install

```bash
go get github.com/sdkwork/sdk-common-go
```

## Authentication Modes

Use one mode per client config.

1. API Key mode
- `Authorization: Bearer {apiKey}`

2. Dual-token mode
- `Access-Token: {accessToken}`
- `Authorization: Bearer {authToken}`

## Quick Start

```go
package main

import (
    "fmt"

    common "github.com/sdkwork/sdk-common-go/common"
)

func main() {
    cfg := common.SdkConfig{
        HttpClientConfig: common.HttpClientConfig{
            BaseURL: "https://api.example.com",
        },
        AccessToken: "your-access-token",
        AuthToken:   "your-auth-token",
    }

    client := common.NewBaseHttpClient(cfg)
    body, err := client.Get("/v1/profile", nil)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(body))
}
```

API key mode example:

```go
cfg := common.SdkConfig{
    HttpClientConfig: common.HttpClientConfig{BaseURL: "https://api.example.com"},
    ApiKey: "your-api-key",
}
```

## Modules

- `common/core`: shared types and constants
- `common/auth`: token manager and auth helpers
- `common/http`: base HTTP client
- `common/errors`: SDK errors
- `common/utils`: retry and cache helpers


## Publishing

This SDK includes cross-platform publish scripts in `bin/`:
- `bin/publish-core.mjs`
- `bin/publish.sh`
- `bin/publish.ps1`

### Check

```bash
./bin/publish.sh --action check
```

### Publish

```bash
./bin/publish.sh --action publish --channel release
```

```powershell
.\bin\publish.ps1 --action publish --channel test --dry-run
```

> Set `GO_RELEASE_TAG` (or `SDKWORK_RELEASE_TAG`) and push tag if needed.

## License

MIT