# SDKwork Java SDK Common

Common Java building blocks for generated SDKs.

## Install

```xml
<dependency>
  <groupId>com.sdkwork</groupId>
  <artifactId>sdk-common</artifactId>
  <version>1.0.0</version>
</dependency>
```

## Current Scope

The Java common module currently provides `core` models and constants:

- `com.sdkwork.common.core.Types`
- `com.sdkwork.common.core.Constants`

## Quick Start

```java
import com.sdkwork.common.core.Constants;
import com.sdkwork.common.core.Types;

Types.SdkConfig config = new Types.SdkConfig("https://api.example.com");
int ok = Constants.HttpStatus.OK;
```

## Authentication Contract

Generated Java SDKs should follow one auth mode per client instance:

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

> Use Maven `settings.xml` credentials and optional `MAVEN_PUBLISH_PROFILE`.

## License

MIT