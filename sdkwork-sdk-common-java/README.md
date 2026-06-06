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

## SDKWork Documentation Contract

Domain: platform
Capability: sdk-common
Package type: java-module
Status: standard

### Public API

Public exports are declared in `specs/component.spec.json` under `contracts.publicExports`.

### Required SDK Surface

- None declared in `specs/component.spec.json`.

### Configuration

Configuration keys and runtime entrypoints are declared in `specs/component.spec.json`.

### SaaS/Private/Local Behavior

This module follows the canonical standards linked from `specs/component.spec.json`, including deployment and runtime configuration rules where applicable.

### Security

Do not add secrets, live tokens, manual auth headers, or app-local credential handling to this module.

### Extension Points

Extension points are limited to declared public exports, runtime entrypoints, SDK clients, events, and config keys.

### Verification

- `mvn test`

### Owner And Status

Owner and lifecycle status are tracked in `specs/component.spec.json`.
