# @sdkwork/sdk-common

Common foundation package for generated TypeScript SDKs.

## Install

```bash
npm install @sdkwork/sdk-common
```

## Authentication Modes

Choose one mode per client instance.

1. API Key mode
- `Authorization: Bearer {apiKey}`

2. Dual-token mode
- `Access-Token: {accessToken}`
- `Authorization: Bearer {authToken}`

## Quick Start

```typescript
import { createBaseHttpClient } from '@sdkwork/sdk-common/http';
import { createTokenManager } from '@sdkwork/sdk-common';

const tokenManager = createTokenManager({
  accessToken: 'your-access-token',
  authToken: 'your-auth-token',
});

const client = createBaseHttpClient({
  baseUrl: 'https://api.example.com',
  tokenManager,
});

const profile = await client.get<{ id: string; name: string }>('/v1/profile');
console.log(profile.name);
```

API key mode example:

```typescript
import { createBaseHttpClient } from '@sdkwork/sdk-common/http';

const client = createBaseHttpClient({
  baseUrl: 'https://api.example.com',
  apiKey: 'your-api-key',
});
```

## Exported Modules

- `core`: request/result types, constants, retry/cache/logger config types
- `auth`: token manager and auth header builder
- `http`: `BaseHttpClient` and `createBaseHttpClient`
- `errors`: SDK error hierarchy and type guards
- `utils`: retry, cache, logger, string/encoding/date/object helpers

## License

MIT