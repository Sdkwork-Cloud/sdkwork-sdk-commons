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

## License

MIT