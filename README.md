# SDKwork SDK Commons

`sdkwork-sdk-commons` contains language-specific common packages used by generated SDKs.

These packages provide the shared building blocks for SDK clients:

- Request/response envelope types
- Authentication header conventions
- HTTP base client abstractions
- Standard error models
- Reusable utility helpers

## Authentication Contract

All SDKs follow one of the two authentication modes below.
The modes are mutually exclusive in configuration.

1. API Key mode
- Header: `Authorization: Bearer {apiKey}`

2. Dual-token mode
- Header: `Access-Token: {accessToken}`
- Header: `Authorization: Bearer {authToken}`

Use one mode per client instance. Do not mix API key and dual-token credentials at the same time.

## Modules By Language

| Language | Directory | Package ID | Scope |
|---|---|---|---|
| TypeScript | `sdkwork-sdk-common-typescript` | `@sdkwork/sdk-common` | `core`, `auth`, `http`, `errors`, `utils` |
| Python | `sdkwork-sdk-common-python` | `sdkwork-common` | `core`, `auth`, `http`, `errors`, `utils` |
| Go | `sdkwork-sdk-common-go` | `github.com/sdkwork/sdk-common-go` | `core`, `auth`, `http`, `errors`, `utils` |
| Java | `sdkwork-sdk-common-java` | `com.sdkwork:sdk-common` | `core` |
| Kotlin | `sdkwork-sdk-common-kt` | `com.sdkwork:sdk-common` | `core` |
| Swift | `sdkwork-sdk-common-swift` | `SDKworkCommon` | `core` |
| Flutter | `sdkwork-sdk-common-flutter` | `sdkwork_common_flutter` | `core`, `http` |
| C# | `sdkwork-sdk-common-csharp` | `SDKwork.Common` | `core`, `errors` |

## Directory Structure

```text
sdkwork-sdk-commons/
  sdkwork-sdk-common-typescript/
  sdkwork-sdk-common-python/
  sdkwork-sdk-common-go/
  sdkwork-sdk-common-java/
  sdkwork-sdk-common-kt/
  sdkwork-sdk-common-swift/
  sdkwork-sdk-common-flutter/
  sdkwork-sdk-common-csharp/
```

## Language READMEs

- [TypeScript](./sdkwork-sdk-common-typescript/README.md)
- [Python](./sdkwork-sdk-common-python/README.md)
- [Go](./sdkwork-sdk-common-go/README.md)
- [Java](./sdkwork-sdk-common-java/README.md)
- [Kotlin](./sdkwork-sdk-common-kt/README.md)
- [Swift](./sdkwork-sdk-common-swift/README.md)
- [Flutter](./sdkwork-sdk-common-flutter/README.md)
- [C#](./sdkwork-sdk-common-csharp/README.md)

## Publish Standard (Per Language `bin`)

Each language directory includes:

- `bin/publish-core.mjs` (shared cross-platform publish core)
- `bin/publish.sh` (Linux/macOS entrypoint)
- `bin/publish.ps1` (Windows/PowerShell entrypoint)

Example:

```bash
cd sdk/sdkwork-sdk-commons/sdkwork-sdk-common-typescript
./bin/publish.sh --action check
./bin/publish.sh --action publish --channel release
```

```powershell
cd sdk/sdkwork-sdk-commons/sdkwork-sdk-common-java
.\bin\publish.ps1 --action publish --channel test --dry-run
```

Target publish ecosystems:

- TypeScript -> npm
- Python -> PyPI/TestPyPI
- Java/Kotlin -> Maven Central compatible publishing
- C# -> NuGet
- Flutter -> pub.dev
- Go/Swift -> git tag based release flow

## Compatibility Notes

- Generated SDKs should consume the matching language common package.
- Keep auth headers aligned with the contract in this document.
- Keep API result and pagination envelope handling consistent across languages.

## License

MIT
