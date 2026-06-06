# sdkwork-common (Python)

Common foundation package for generated Python SDKs.

## Install

```bash
pip install sdkwork-common
```

## Authentication Modes

Choose one mode when creating a client config.

1. API Key mode
- `Authorization: Bearer {apiKey}`

2. Dual-token mode
- `Access-Token: {accessToken}`
- `Authorization: Bearer {authToken}`

## Quick Start

```python
from sdkwork.common.core.types import SdkConfig
from sdkwork.common.http import create_http_client

config = SdkConfig(
    base_url="https://api.example.com",
    access_token="your-access-token",
    auth_token="your-auth-token",
)

client = create_http_client(config)
result = client.get("/v1/profile")
print(result)
```

API key mode example:

```python
from sdkwork.common.core.types import SdkConfig

config = SdkConfig(
    base_url="https://api.example.com",
    api_key="your-api-key",
)
```

## Modules

- `core`: common types and constants
- `auth`: token objects and token manager helpers
- `http`: base HTTP client
- `errors`: SDK error types and helpers
- `utils`: retry/cache/logger helpers


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

> Set `PYPI_TOKEN` for release (or `TEST_PYPI_TOKEN` for test channel).

## License

MIT

## SDKWork Documentation Contract

Domain: platform
Capability: common
Package type: python-package
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

- `powershell -NoProfile -Command "Get-Content specs/component.spec.json -Raw | ConvertFrom-Json | Out-Null"`

### Owner And Status

Owner and lifecycle status are tracked in `specs/component.spec.json`.
