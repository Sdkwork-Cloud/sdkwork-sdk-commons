# sdkwork_common_flutter

Common Flutter/Dart foundation package for generated SDKs.

## Install

```yaml
dependencies:
  sdkwork_common_flutter: ^1.0.0
```

## Authentication Modes

Use one mode per client instance.

1. API Key mode
- `Authorization: Bearer {apiKey}` (or custom API key header)

2. Dual-token mode
- `Access-Token: {accessToken}`
- `Authorization: Bearer {authToken}`

## Quick Start

```dart
import 'package:sdkwork_common_flutter/sdkwork_common_flutter.dart';

final client = BaseHttpClient(
  const SdkConfig(
    baseUrl: 'https://api.example.com',
    accessToken: 'your-access-token',
    authToken: 'your-auth-token',
  ),
);

final profile = await client.get('/v1/profile');
print(profile);
```

API key mode example:

```dart
final client = BaseHttpClient(
  const SdkConfig(
    baseUrl: 'https://api.example.com',
    apiKey: 'your-api-key',
  ),
);
```

## Exports

- `SdkConfig`
- `BaseHttpClient`


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

> Ensure `dart pub publish --dry-run` passes before release publish.

## License

MIT