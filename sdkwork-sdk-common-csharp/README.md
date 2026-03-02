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

## License

MIT