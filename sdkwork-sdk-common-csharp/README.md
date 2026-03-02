# SDKwork C# SDK Common

Common C# (.NET) building blocks for generated SDKs.

## Install

```bash
dotnet add package SDKwork.Common
```

## Current Scope

The C# common module currently provides:

- `SDKwork.Common.Core`: shared types and constants
- `SDKwork.Common.Errors`: SDK exception hierarchy and helpers

## Quick Start

```csharp
using SDKwork.Common.Core;
using SDKwork.Common.Errors;

var config = new SdkConfig(BaseUrl: "https://api.example.com");
var isOk = HttpStatus.OK == 200;

try
{
    throw new AuthenticationException("Invalid token");
}
catch (SdkException ex)
{
    Console.WriteLine($"Code={ex.Code}, Message={ex.Message}");
}
```

## Authentication Contract

Generated C# SDKs should follow one auth mode per client instance:

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

> Set `NUGET_API_KEY` for release (or `NUGET_TEST_API_KEY` for test channel).

## License

MIT