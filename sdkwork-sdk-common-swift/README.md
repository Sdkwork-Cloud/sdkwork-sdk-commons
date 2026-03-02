# SDKwork Swift SDK Common

Common Swift building blocks for generated SDKs.

## Install

Add dependency in `Package.swift`:

```swift
dependencies: [
    .package(url: "https://github.com/sdkwork/sdk-common-swift", from: "1.0.0")
]
```

## Current Scope

The Swift common module currently provides `Core` data types and constants:

- `SdkConfig`
- `ApiResult`
- `PageResult`
- `HttpStatus`
- `MimeTypes`

## Quick Start

```swift
import SDKworkCommon

let config = SdkConfig(
    baseUrl: "https://api.example.com",
    authToken: "your-auth-token",
    accessToken: "your-access-token"
)

print(HttpStatus.ok)
```

## Authentication Contract

Use one mode per client instance:

1. API Key mode: `Authorization: Bearer {apiKey}`
2. Dual-token mode: `Access-Token: {accessToken}` and `Authorization: Bearer {authToken}`


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

> Set `SWIFT_RELEASE_TAG` (or `SDKWORK_RELEASE_TAG`) for tag-based release.

## License

MIT