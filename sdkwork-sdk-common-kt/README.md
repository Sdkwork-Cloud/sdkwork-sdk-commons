# SDKwork Kotlin SDK Common

Common Kotlin building blocks for generated SDKs.

## Install

```kotlin
dependencies {
    implementation("com.sdkwork:sdk-common:1.0.0")
}
```

## Current Scope

The Kotlin common module currently provides `core` models and constants:

- `com.sdkwork.common.core.SdkConfig`
- `com.sdkwork.common.core.ApiResult`
- `com.sdkwork.common.core.PageResult`
- `com.sdkwork.common.core.HttpStatus`

## Quick Start

```kotlin
import com.sdkwork.common.core.HttpStatus
import com.sdkwork.common.core.SdkConfig

val config = SdkConfig(baseUrl = "https://api.example.com")
val ok = HttpStatus.OK
```

## Authentication Contract

Generated Kotlin SDKs should follow one auth mode per client instance:

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

> Configure Gradle publishing credentials and optional `GRADLE_PUBLISH_TASK`.

## License

MIT