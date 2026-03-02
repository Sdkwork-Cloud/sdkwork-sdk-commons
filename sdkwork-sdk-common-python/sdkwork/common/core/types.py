from typing import Any, Optional, Dict, List, Callable, Union, Literal
from dataclasses import dataclass, field
from enum import Enum
import time
import threading

HttpMethod = Literal['GET', 'POST', 'PUT', 'DELETE', 'PATCH', 'HEAD', 'OPTIONS']
LogLevel = Literal['debug', 'info', 'warn', 'error', 'silent']
QueryParams = Dict[str, Optional[Union[str, int, float, bool]]]
HttpHeaders = Dict[str, str]


@dataclass
class ApiResult:
    code: Union[int, str]
    data: Any = None
    msg: Optional[str] = None
    message: Optional[str] = None
    timestamp: Optional[int] = None
    trace_id: Optional[str] = None


@dataclass
class PageResult:
    content: Optional[List[Any]] = None
    list: Optional[List[Any]] = None
    total: int = 0
    total_elements: Optional[int] = None
    page: int = 1
    page_size: int = 10
    size: Optional[int] = None
    total_pages: int = 0
    has_more: bool = False
    first: Optional[bool] = None
    last: Optional[bool] = None
    empty: Optional[bool] = None
    number: Optional[int] = None


@dataclass
class Pageable:
    page: Optional[int] = 1
    page_size: Optional[int] = 10
    size: Optional[int] = None
    sort: Optional[str] = None
    order: Optional[str] = None


class RetryBackoff(str, Enum):
    FIXED = 'fixed'
    LINEAR = 'linear'
    EXPONENTIAL = 'exponential'


@dataclass
class RetryConfig:
    max_retries: int = 3
    retry_delay: float = 1.0
    retry_backoff: RetryBackoff = RetryBackoff.EXPONENTIAL
    max_retry_delay: float = 30.0
    retry_condition: Optional[Callable[[Exception, int], bool]] = None


@dataclass
class CacheConfig:
    enabled: bool = False
    ttl: int = 300000
    max_size: int = 100


@dataclass
class LoggerConfig:
    level: LogLevel = 'info'
    prefix: str = '[SDK]'
    timestamp: bool = True
    colors: bool = True


@dataclass
class RequestConfig:
    url: str
    method: HttpMethod = 'GET'
    headers: Optional[HttpHeaders] = None
    params: Optional[QueryParams] = None
    body: Optional[Any] = None
    timeout: Optional[int] = None
    skip_auth: bool = False
    retry_count: int = 0
    metadata: Optional[Dict[str, Any]] = None


@dataclass
class HttpClientConfig:
    base_url: str = ''
    timeout: Optional[int] = None
    headers: Optional[HttpHeaders] = None
    retry: Optional[RetryConfig] = None
    cache: Optional[CacheConfig] = None
    logger: Optional[LoggerConfig] = None


@dataclass
class SdkConfig(HttpClientConfig):
    api_key: Optional[str] = None
    auth_token: Optional[str] = None
    access_token: Optional[str] = None
    tenant_id: Optional[str] = None
    organization_id: Optional[str] = None
    platform: Optional[str] = None
    user_id: Optional[Union[str, int]] = None


DEFAULT_RETRY_CONFIG = RetryConfig()
DEFAULT_CACHE_CONFIG = CacheConfig()
DEFAULT_LOGGER_CONFIG = LoggerConfig()
DEFAULT_TIMEOUT = 30000

SUCCESS_CODES = [0, 200, 2000, '0', '200', '2000']

HTTP_STATUS = {
    'OK': 200,
    'CREATED': 201,
    'NO_CONTENT': 204,
    'BAD_REQUEST': 400,
    'UNAUTHORIZED': 401,
    'FORBIDDEN': 403,
    'NOT_FOUND': 404,
    'METHOD_NOT_ALLOWED': 405,
    'CONFLICT': 409,
    'UNPROCESSABLE_ENTITY': 422,
    'TOO_MANY_REQUESTS': 429,
    'INTERNAL_SERVER_ERROR': 500,
    'BAD_GATEWAY': 502,
    'SERVICE_UNAVAILABLE': 503,
    'GATEWAY_TIMEOUT': 504,
}

MIME_TYPES = {
    'JSON': 'application/json',
    'FORM_DATA': 'multipart/form-data',
    'URL_ENCODED': 'application/x-www-form-urlencoded',
    'OCTET_STREAM': 'application/octet-stream',
    'TEXT_PLAIN': 'text/plain',
    'TEXT_HTML': 'text/html',
}

__all__ = [
    'HttpMethod',
    'LogLevel',
    'QueryParams',
    'HttpHeaders',
    'ApiResult',
    'PageResult',
    'Pageable',
    'RetryBackoff',
    'RetryConfig',
    'CacheConfig',
    'LoggerConfig',
    'RequestConfig',
    'HttpClientConfig',
    'SdkConfig',
    'DEFAULT_RETRY_CONFIG',
    'DEFAULT_CACHE_CONFIG',
    'DEFAULT_LOGGER_CONFIG',
    'DEFAULT_TIMEOUT',
    'SUCCESS_CODES',
    'HTTP_STATUS',
    'MIME_TYPES',
]
