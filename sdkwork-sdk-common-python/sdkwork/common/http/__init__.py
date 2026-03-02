from typing import Any, Optional, Dict, List, Callable
from dataclasses import dataclass, field
from datetime import datetime, timedelta
from enum import Enum
import time
import logging

from ..core.types import (
    SdkConfig,
    HttpClientConfig,
    RetryConfig,
    CacheConfig,
    LoggerConfig,
    HttpHeaders,
    QueryParams,
    DEFAULT_RETRY_CONFIG,
    DEFAULT_CACHE_CONFIG,
    DEFAULT_LOGGER_CONFIG,
    HTTP_STATUS,
)
from ..errors import (
    SdkError,
    NetworkError,
    TimeoutError,
    AuthenticationError,
    RateLimitError,
    ServerError,
    ForbiddenError,
    NotFoundError,
    ValidationError,
    is_retryable_error,
)


logger = logging.getLogger('sdkwork')


class BaseHttpClient:
    def __init__(self, config: SdkConfig):
        self.config = config
        self.base_url = config.base_url.rstrip('/')
        self.timeout = config.timeout or 30000
        self.headers = config.headers or {}
        self._api_key = config.api_key
        self._auth_token = config.auth_token
        self._access_token = config.access_token
        self._session = None

    def _get_session(self):
        import requests
        if self._session is None:
            self._session = requests.Session()
            self._session.headers.update(self.headers)
            self._update_auth_headers()
        return self._session

    def _update_auth_headers(self):
        if self._session is None:
            return
        self._session.headers.pop('Authorization', None)
        self._session.headers.pop('Access-Token', None)
        self._session.headers.pop('X-API-Key', None)

        if self._api_key:
            self._session.headers['Authorization'] = f'Bearer {self._api_key}'
        if self._auth_token:
            self._session.headers['Authorization'] = f'Bearer {self._auth_token}'
        if self._access_token:
            self._session.headers['Access-Token'] = self._access_token

    def set_api_key(self, api_key: str) -> 'BaseHttpClient':
        self._api_key = api_key
        self._update_auth_headers()
        return self

    def set_auth_token(self, token: str) -> 'BaseHttpClient':
        self._auth_token = token
        self._update_auth_headers()
        return self

    def set_access_token(self, token: str) -> 'BaseHttpClient':
        self._access_token = token
        self._update_auth_headers()
        return self

    def request(
        self,
        method: str,
        path: str,
        params: Optional[QueryParams] = None,
        data: Any = None,
        json: Any = None,
        headers: Optional[HttpHeaders] = None,
        timeout: Optional[int] = None,
        **kwargs
    ) -> Any:
        import requests
        url = self.base_url + path
        request_headers = {**self.headers, **(headers or {})}
        request_timeout = (timeout or self.timeout) / 1000

        try:
            response = self._get_session().request(
                method=method,
                url=url,
                params=params,
                data=data,
                json=json,
                headers=request_headers,
                timeout=request_timeout,
                **kwargs
            )
            response.raise_for_status()
            return response.json() if response.content else None
        except requests.exceptions.Timeout as e:
            raise TimeoutError(str(e), timeout=int(request_timeout * 1000))
        except requests.exceptions.ConnectionError as e:
            raise NetworkError(str(e))
        except requests.exceptions.HTTPError as e:
            raise self._handle_http_error(e)

    def get(self, path: str, params: Optional[QueryParams] = None, **kwargs) -> Any:
        return self.request('GET', path, params=params, **kwargs)

    def post(self, path: str, data: Any = None, json: Any = None, **kwargs) -> Any:
        return self.request('POST', path, data=data, json=json, **kwargs)

    def put(self, path: str, data: Any = None, json: Any = None, **kwargs) -> Any:
        return self.request('PUT', path, data=data, json=json, **kwargs)

    def delete(self, path: str, **kwargs) -> Any:
        return self.request('DELETE', path, **kwargs)

    def patch(self, path: str, data: Any = None, json: Any = None, **kwargs) -> Any:
        return self.request('PATCH', path, data=data, json=json, **kwargs)

    def _handle_http_error(self, error) -> SdkError:
        import requests
        status = error.response.status_code
        try:
            data = error.response.json()
            message = data.get('msg') or data.get('message') or str(error)
        except:
            message = str(error)

        if status == HTTP_STATUS['UNAUTHORIZED']:
            return AuthenticationError(message)
        elif status == HTTP_STATUS['FORBIDDEN']:
            return ForbiddenError(message)
        elif status == HTTP_STATUS['NOT_FOUND']:
            return NotFoundError(message)
        elif status == HTTP_STATUS['BAD_REQUEST']:
            return ValidationError(message)
        elif status == HTTP_STATUS['TOO_MANY_REQUESTS']:
            retry_after = error.response.headers.get('Retry-After')
            return RateLimitError(message, retry_after=int(retry_after) if retry_after else None)
        elif status >= HTTP_STATUS['INTERNAL_SERVER_ERROR']:
            return ServerError(message, status)
        else:
            return SdkError(message, http_status=status)


def create_http_client(config: SdkConfig) -> BaseHttpClient:
    return BaseHttpClient(config)


__all__ = ['BaseHttpClient', 'create_http_client']
